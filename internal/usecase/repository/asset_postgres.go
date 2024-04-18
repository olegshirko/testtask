package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"testTask/internal/entity"
	"testTask/pkg/postgres"
)

const _defaultEntityCap = 64

// AssetRepo -.
type AssetRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *AssetRepo {
	return &AssetRepo{pg}
}

// GetUserByAuthData -.
func (r *AssetRepo) GetUserByAuthData(ctx context.Context, authdata entity.AuthData) (entity.User, error) {
	sql := "select * from users where login = @login limit 1;"

	args := pgx.NamedArgs{
		"login": authdata.Login,
	}

	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return entity.User{}, fmt.Errorf("AssetRepo - GetUserByAuthData - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, nil
		}
		return entity.User{}, fmt.Errorf("AssetRepo - GetUserByAuthData - r.Pool.Exec: %w", err)
	}
	return user, nil
}

// GetSession -.
func (r *AssetRepo) GetSession(ctx context.Context, user entity.User) (entity.Session, error) {
	sql := "select id, uid, created_at, user_ip::text from sessions " +
		"where uid = @id order by created_at desc limit 1;"

	args := pgx.NamedArgs{
		"id": user.Id,
	}

	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AssetRepo - GetSession - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	session, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Session])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Session{}, nil
		}
		return entity.Session{}, fmt.Errorf("AssetRepo - GetSession - r.Pool.Exec: %w", err)
	}
	return session, nil
}

// StoreSession -.
func (r *AssetRepo) StoreSession(ctx context.Context, user entity.User, ip string) (entity.Session, error) {
	sql := "insert into sessions (uid, user_ip) values (@uid, @ip::inet) returning id, uid, created_at, user_ip::text;"

	args := pgx.NamedArgs{
		"uid": user.Id,
		"ip":  ip,
	}
	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return entity.Session{}, fmt.Errorf("AssetRepo - StoreSession - r.Pool.Exec: %w", err)
	}
	defer rows.Close()

	session, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Session])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Session{}, nil
		}
		return entity.Session{}, fmt.Errorf("AssetRepo - StoreSession - r.Pool.Exec: %w", err)
	}
	return session, nil
}

// StoreAsset -.
func (r *AssetRepo) StoreAsset(ctx context.Context, data []byte, name string, uid int64) error {
	sql := "insert into assets (name, uid, data) values (@name, @uid, @data);"

	args := pgx.NamedArgs{
		"name": name,
		"uid":  uid,
		"data": data,
	}
	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return fmt.Errorf("AssetRepo - StoreAsset - r.Pool.Exec: %w", err)
	}
	defer rows.Close()

	return nil
}

// GetHistory -.
func (r *AssetRepo) GetHistory(ctx context.Context, uid int64) ([]entity.Asset, error) {
	sql := "select * from assets where uid = @uid order by created_at;"

	args := pgx.NamedArgs{
		"uid": uid,
	}

	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return nil, fmt.Errorf("AssetRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	assets := make([]entity.Asset, 0, _defaultEntityCap)

	for rows.Next() {
		a := entity.Asset{}

		err = rows.Scan(&a.Name, &a.Uid, &a.Data, &a.Created)
		if err != nil {
			return nil, fmt.Errorf("AssetRepo - GetHistory - rows.Scan: %w", err)
		}

		assets = append(assets, a)
	}

	return assets, nil
}

// GetUserIdByToken -.
func (r *AssetRepo) GetUserIdByToken(ctx context.Context, token string) (int64, error) {
	sql := "select uid from sessions where id = @token order by created_at desc limit 1;"

	args := pgx.NamedArgs{
		"token": token,
	}

	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return 0, fmt.Errorf("AssetRepo - GetUserIdByToken - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	uid, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("AssetRepo - GetUserIdByToken - r.Pool.Exec: %w", err)
	}
	return uid, nil
}

// DeleteAsset
func (r *AssetRepo) DeleteAsset(ctx context.Context, name string, uid int64) (int64, error) {
	sql := "WITH deleted AS (delete from assets where name = @name and uid = @uid returning *) " +
		"SELECT count(*) FROM deleted;"

	args := pgx.NamedArgs{
		"uid":  uid,
		"name": name,
	}

	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return 0, fmt.Errorf("AssetRepo - DeleteAsset - r.Pool.Exec: %w", err)
	}
	defer rows.Close()

	cnt, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])

	if err != nil {
		return 0, fmt.Errorf("AssetRepo - DeleteAsset - r.Pool.Exec: %w", err)
	}

	return cnt, nil

}
