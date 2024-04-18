package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net"
	"net/http"
	"strings"
	"testTask/internal/entity"
	"testTask/internal/usecase"
	"testTask/pkg/logger"
	"time"
)

type assetRoutes struct {
	t usecase.Asset
	l logger.Interface
}

type UserCtxKey struct{}

func newAssetRoutes(r chi.Router, t usecase.Asset, l logger.Interface) {
	// Create an instance of assetRoutes that holds dependencies.
	ar := &assetRoutes{t, l}

	// Setting up routes with the appropriate methods and route patterns.

	r.Post("/auth", ar.doAuth)
	r.With(ar.SessionIDChecker).Get("/history", ar.history)
	r.With(ar.SessionIDChecker).Post("/upload-asset/{assetName}", ar.doUpload)
	r.With(ar.SessionIDChecker).Delete("/del-asset/{assetName}", ar.doDelete)
}

func (ar *assetRoutes) SessionIDChecker(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := extractBearerToken(r)
		if err != nil {
			ErrorResponse(w, http.StatusUnauthorized, "Unauthorized: No valid bearer token provided")
			return
		}

		// Call the use case to handle the history logic.
		uid, err := ar.t.GetUserIdByToken(r.Context(), token)
		if err != nil {
			ar.l.Error(err, "http - v1 - SessionIDChecker")
			ErrorResponse(w, http.StatusUnauthorized, "Unauthorized: No valid bearer token provided")
			return
		}

		//add user id to context
		ctx := context.WithValue(r.Context(), UserCtxKey{}, uid)
		// If valid, proceed with the next handler
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

// @Summary     Auth
// @Description Auth a user
// @Id          auth
// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param       request body doAuthResponse true "request body"
// @Success     200 {object} doAuthResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /auth [post]
func (ar *assetRoutes) doAuth(w http.ResponseWriter, r *http.Request) {
	var request doAuthRequest
	// Parse the JSON body into the request struct
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		ar.l.Error(err, "http - v1 - doAuth")
		ErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Extract the IP address from the request
	ip := getIP(r)

	// Validate required fields
	if !isValidLogin(request.Login) || !isValidPassword(request.Password) {
		ErrorResponse(w, http.StatusBadRequest, "invalid login or password format")
		return
	}

	// Call the underlying logic to authenticate the user.
	session, err := ar.t.Session(r.Context(), ip, entity.AuthData{
		Login:    request.Login,
		Password: request.Password,
	})

	// Handle possible authentication errors
	if err != nil {
		ar.l.Error(err, "http - v1 - doAuth")
		ErrorResponse(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// Successful authentication
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doAuthResponse{session.Id})
}

type doAuthRequest struct {
	Login    string `json:"login"       binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type doAuthResponse struct {
	History string `json:"token"`
}

// @Summary     Show history
// @Description Show all file upload history
// @Id          history
// @Tags  	    asset
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /asset/history [get]
func (ar *assetRoutes) history(w http.ResponseWriter, r *http.Request) {
	uid, ok := getUserIDFromContext(r.Context())
	if !ok || uid == 0 {
		errorResponse(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Call the underlying logic to retrieve asset history.
	assets, err := ar.t.History(r.Context(), uid)
	if err != nil {
		ar.l.Error(err, "http - v1 - history")
		errorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	// Construct the response struct.
	var respAssets []historyResponseAsset
	for _, a := range assets {
		respAssets = append(respAssets, historyResponseAsset{
			AssetName: a.Name,
			Created:   a.Created,
		})
	}

	// Encode the response as JSON and send it.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historyResponse{respAssets})
}

type historyResponse struct {
	History []historyResponseAsset `json:"History"`
}

type historyResponseAsset struct {
	AssetName string
	Created   time.Time
}

// @Summary     Upload
// @Description Upload a data
// @Id          do-upload
// @Tags  	    upload
// @Accept      json
// @Produce     json
// @Success     200 {object} doUploadResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /upload-asset/:assetName [post]
func (ar *assetRoutes) doUpload(w http.ResponseWriter, r *http.Request) {
	uid, ok := getUserIDFromContext(r.Context())
	if !ok || uid == 0 {
		ErrorResponse(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Check if the content length is 0 which means no data is sent.
	if r.ContentLength <= 0 {
		ErrorResponse(w, http.StatusBadRequest, "No data in request body")
		return
	}

	// Define a maximum size for the upload
	const maxUploadSize = 10 * 1024 * 1024 // 10 MB

	// Check if the request body size exceeds the maximum limit.
	if r.ContentLength > maxUploadSize {
		ErrorResponse(w, http.StatusRequestEntityTooLarge, "Upload exceeds maximum limit of 10 MB")
		return
	}

	// Read the raw data from the request body.
	data, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		ar.l.Error(err, "http - v1 - doUpload - failed to read raw data")
		ErrorResponse(w, http.StatusInternalServerError, "Failed to read raw data")
		return
	}

	// Extract the asset name from the URL parameter.
	assetName := chi.URLParam(r, "assetName")

	// Call the use case to handle the upload logic.
	err = ar.t.UploadAsset(r.Context(), data, assetName, uid)
	if err != nil {
		ar.l.Error(err, "http - v1 - doUpload")
		ErrorResponse(w, http.StatusInternalServerError, "Asset service problems")
		return
	}

	// If successful, send a success response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doUploadResponse{"ok"})
}

type doUploadResponse struct {
	Status string `json:"Status"`
}

// @Summary     Delete asset
// @Description Delete asset
// @Id          del-asset
// @Tags  	    asset
// @Accept      json
// @Produce     json
// @Success     200 {object} doDeleteResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /del-asset/:assetName [delete]
func (ar *assetRoutes) doDelete(w http.ResponseWriter, r *http.Request) {
	uid, ok := getUserIDFromContext(r.Context())
	if !ok || uid == 0 {
		ErrorResponse(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Extract the asset name from the URL parameter.
	assetName := chi.URLParam(r, "assetName")

	// Call the use case to handle the upload logic.
	cnt, err := ar.t.DropAsset(r.Context(), assetName, uid)
	if err != nil || cnt == 0 {
		ar.l.Error(err, "http - v1 - doDelete")
		ErrorResponse(w, http.StatusInternalServerError, "Asset not found for delete")
		return
	}

	// If successful, send a success response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doDeleteResponse{fmt.Sprintf("%d asset deleted", cnt)})

}

type doDeleteResponse struct {
	History string `json:"Status"`
}

func extractBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", http.ErrMissingBoundary
	}

	// Split the header into "Bearer" and the <token>
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", http.ErrHeaderTooLong
	}

	return parts[1], nil
}

func getUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserCtxKey{}).(int64)
	return userID, ok
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{Error: message})
}

type Response struct {
	Error string `json:"error"`
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		for _, ip := range ips {
			trimmedIP := strings.TrimSpace(ip)
			if trimmedIP != "" {
				return trimmedIP
			}
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func isValidLogin(login string) bool {
	return login != ""
}

func isValidPassword(password string) bool {
	return len(password) >= 6
}
