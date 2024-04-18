package usecase

import (
	"context"
	"testTask/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Asset -.
	Asset interface {
		History(context.Context, int64) ([]entity.Asset, error)
		Session(context.Context, string, entity.AuthData) (entity.Session, error)
		UploadAsset(context.Context, []byte, string, int64) error
		GetUserIdByToken(context.Context, string) (int64, error)
		DropAsset(context.Context, string, int64) (int64, error)
	}

	// AssetRepo -.
	AssetRepo interface {
		//Store(context.Context, entity.Asset) error
		GetHistory(context.Context, int64) ([]entity.Asset, error)
		GetUserByAuthData(context.Context, entity.AuthData) (entity.User, error)
		GetSession(context.Context, entity.User) (entity.Session, error)
		StoreSession(context.Context, entity.User, string) (entity.Session, error)
		StoreAsset(context.Context, []byte, string, int64) error
		GetUserIdByToken(context.Context, string) (int64, error)
		DeleteAsset(context.Context, string, int64) (int64, error)
	}
)
