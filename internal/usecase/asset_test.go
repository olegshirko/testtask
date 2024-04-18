package usecase_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testTask/internal/entity"
	"testTask/internal/usecase"
	"testing"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func asset(t *testing.T) (*usecase.AssetUseCase, *MockAssetRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockAssetRepo(mockCtl)

	asset := usecase.New(repo)

	return asset, repo
}

func TestHistory(t *testing.T) {
	t.Parallel()

	asset, repo := asset(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background(), int64(1)).Return(nil, nil)
			},
			res: []entity.Asset(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background(), int64(1)).Return(nil, errInternalServErr)
			},
			res: []entity.Asset(nil),
			err: errInternalServErr,
		},
		{
			name: "result success",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background(), int64(1)).Return([]entity.Asset{
					{
						Uid:  1,
						Data: nil,
						Name: "hello",
					},
					{
						Uid:  2,
						Data: make([]byte, 1),
					},
				}, nil)
			},
			res: []entity.Asset{{
				Uid:  1,
				Data: nil,
				Name: "hello",
			},
				{
					Uid:  2,
					Data: make([]byte, 1),
				}},
			err: nil,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := asset.History(context.Background(), int64(1))

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}

}

func TestAuth(t *testing.T) {
	t.Parallel()

	asset, repo := asset(t)

	tests := []test{
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetSession(context.Background(), entity.User{}).Return(entity.Session{}, nil)
				repo.EXPECT().GetUserByAuthData(context.Background(), entity.AuthData{}).Return(entity.User{}, errInternalServErr)
			},
			res: entity.Session{},
			err: errInternalServErr,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := asset.Session(context.Background(), string("x"), entity.AuthData{})

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}

}
