package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testTask/internal/entity"
	"time"
)

// AssetUseCase -.
type AssetUseCase struct {
	repo AssetRepo
}

// New -.
func New(r AssetRepo) *AssetUseCase {
	return &AssetUseCase{
		repo: r,
	}
}

func (uc *AssetUseCase) getUser(ctx context.Context, authdata entity.AuthData) (entity.User, error) {
	user, err := uc.repo.GetUserByAuthData(ctx, authdata)
	if err != nil {
		return entity.User{}, fmt.Errorf("AssetUseCase - getUser - s.repo.getUser: %w", err)
	}
	return user, nil
}

func verifyPassword(user entity.User, auth string) bool {
	hashedPassword := hashPassword(auth)
	return hashedPassword == user.Password
}

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func (uc *AssetUseCase) getSession(ctx context.Context, user entity.User) (entity.Session, error) {
	activeSession, err := uc.repo.GetSession(ctx, user)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AssetUseCase - Session - s.repo.getSession: %w", err)
	}
	//if created less 24 h before now then true
	if time.Since(activeSession.Created).Hours() > 24 {
		return entity.Session{}, nil
	}
	return activeSession, nil
}

func (uc *AssetUseCase) Session(ctx context.Context, ip string, authdata entity.AuthData) (entity.Session, error) {
	//get user from repo and check pass
	user, err := uc.getUser(ctx, authdata)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AssetUseCase - Session - uc.getUser: %w", err)
	}

	//check user pass
	if !verifyPassword(user, authdata.Password) {
		return entity.Session{}, fmt.Errorf("AssetUseCase - Session - verifyPassword: %w", err)
	}

	activeSession, err := uc.getSession(ctx, user)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AssetUseCase - Session - s.repo.getSession: %w", err)
	}
	//TODO align to happy path
	if activeSession != (entity.Session{}) {
		return activeSession, nil
	}

	//create new session
	newSession, err := uc.repo.StoreSession(ctx, user, ip)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AssetUseCase - Session - s.repo.StoreSession: %w", err)
	}
	return newSession, nil
}

// History - getting asset upload history from store.
func (uc *AssetUseCase) History(ctx context.Context, uid int64) ([]entity.Asset, error) {
	assets, err := uc.repo.GetHistory(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("AssetUseCase - History - s.repo.GetHistory: %w", err)
	}

	return assets, nil
}

// UploadAsset - save asset to store.
func (uc *AssetUseCase) UploadAsset(ctx context.Context, data []byte, name string, uid int64) error {
	err := uc.repo.StoreAsset(ctx, data, name, uid)
	if err != nil {
		return fmt.Errorf("AssetUseCase - History - s.repo.GetHistory: %w", err)
	}

	return nil
}

// GetUserIdByToken
func (uc *AssetUseCase) GetUserIdByToken(ctx context.Context, token string) (int64, error) {
	//get user id by token
	uid, err := uc.repo.GetUserIdByToken(ctx, token)
	if err != nil || uid == 0 {
		return 0, fmt.Errorf("AssetUseCase - GetUserIdByToken - s.repo.GetUserIdByToken: %w", err)
	}

	//get active session
	activeSession, err := uc.getSession(ctx, entity.User{Id: uid})
	if err != nil {
		return 0, fmt.Errorf("AssetUseCase - GetUserIdByToken - s.repo.GetSession: %w", err)
	}

	//if token is not equal to session token then return 0
	if activeSession.Id != token {
		return 0, fmt.Errorf("AssetUseCase - GetUserIdByToken - s.repo.GetSession: %w", err)

	}

	return uid, nil
}

// DropAsset
func (uc *AssetUseCase) DropAsset(ctx context.Context, name string, uid int64) (int64, error) {
	cnt, err := uc.repo.DeleteAsset(ctx, name, uid)
	if err != nil {
		return 0, fmt.Errorf("AssetUseCase - DropAsset - s.repo.DropAsset: %w", err)
	}

	return cnt, nil

}
