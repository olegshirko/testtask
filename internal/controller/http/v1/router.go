package v1

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"testTask/internal/usecase"
	"testTask/pkg/logger"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"github.com/swaggo/http-swagger/v2"
	// Swagger docs.
	//_ "testTask/docs"
)

// NewRouter -.
// Swagger spec:
// @title       File manage service API
// @description Using a file manage service
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(l logger.Interface, t usecase.Asset, handler *chi.Mux) *chi.Mux {

	// Middlewares
	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)

	// Swagger
	if os.Getenv("DISABLE_SWAGGER_HTTP_HANDLER") == "false" {
		handler.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("doc.json"), // The url pointing to API definition
		))
	}

	// K8s probe
	handler.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Prometheus metrics
	handler.Get("/metrics", promhttp.Handler().ServeHTTP)

	// Routers
	handler.Route("/api", func(r chi.Router) {
		newAssetRoutes(r, t, l)
	})

	return handler
}
