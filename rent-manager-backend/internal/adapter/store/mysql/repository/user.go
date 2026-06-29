package repository

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"log/slog"

	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
)

type UserRepository struct {
	db *mysql.DB
}

func NewUserRepository(db *mysql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := ur.db.QueryBuilder.
		Insert("users").
		Columns(
			"rol_id",
			"usr_name",
			"usr_email",
			"usr_password_hash",
		).
		Values(
			user.RoleID,
			user.Name,
			user.Email,
			user.PasswordHash,
		)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[CreateUser] ToSql error:", err)
		return nil, err
	}

	result, errExe := ur.db.ExecContext(ctx, sqlQuery, args...)
	if errExe != nil {
		slog.Error("[CreateUser] Exec err:", errExe)
		if errCode := ur.db.ErrorCode(errExe); errCode == "1062" {
			return nil, domain.ErrConflictingData
		}
		return nil, errExe
	}

	lastID, errLastInsert := result.LastInsertId()
	if errLastInsert != nil {
		slog.Error("[CreateUser] LastInsertId err:", errLastInsert)
		return nil, errLastInsert
	}

	return ur.GetUserByID(ctx, lastID)
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User

	query := ur.baseUserQuery().
		Where(sq.Eq{"usr_id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetUserByID] Error while building sql query", err)
		return nil, err
	}

	err = ur.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&user.ID,
		&user.RoleID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		slog.Error("[GetUserByID] Error:", err, "id:", id)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	query := ur.baseUserQuery().
		Where(sq.Expr("LOWER(usr_email) = LOWER(?)", email)).
		Limit(1)

	sqlQuery, args, errSql := query.ToSql()
	if errSql != nil {
		slog.Error("[GetUserByEmail] errSql", errSql)
		return nil, errSql
	}

	errQuery := ur.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&user.ID,
		&user.RoleID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if errQuery != nil {
		slog.Error("[GetUserByEmail] Error:", errQuery)
		if errors.Is(errQuery, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}
		return nil, errQuery
	}

	return &user, nil
}

func (ur *UserRepository) ListUsers(ctx context.Context, page uint64, limit uint64) ([]domain.User, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := ur.baseUserQuery().
		OrderBy("usr_id").
		Limit(limit).
		Offset(offset)

	return ur.findUsers(ctx, query)
}

func (ur *UserRepository) SearchUsersByNameOrEmail(
	ctx context.Context,
	value string,
	page uint64,
	limit uint64) ([]domain.User, error) {

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	searchValue := "%" + value + "%"

	query := ur.baseUserQuery().
		Where(
			sq.Or{
				sq.Expr("LOWER(usr_name) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(usr_email) LIKE LOWER(?)", searchValue),
			},
		).
		OrderBy("usr_id").
		Limit(limit).
		Offset(offset)

	return ur.findUsers(ctx, query)
}

func (ur *UserRepository) PathUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	queryBuilder := ur.db.QueryBuilder.Update("users")

	if user.RoleID != 0 {
		queryBuilder = queryBuilder.Set("rol_id", user.RoleID)
	}

	if user.Name != "" {
		queryBuilder = queryBuilder.Set("usr_name", user.Name)
	}

	if user.Email != "" {
		queryBuilder = queryBuilder.Set("usr_email", user.Email)
	}

	if user.PasswordHash != "" {
		queryBuilder = queryBuilder.Set("usr_password_hash", user.PasswordHash)
	}
	queryBuilder = queryBuilder.Where(sq.Eq{"usr_id": user.ID})
	sqlQuery, args, err := queryBuilder.ToSql()
	slog.Info("[PathUser] executing user update query", "sql", sqlQuery, "args", args)
	if err != nil {
		slog.Error("[PathUser] failed to generate sql query", "error", err)
		return nil, err
	}
	result, errExe := ur.db.ExecContext(ctx, sqlQuery, args...)
	if errExe != nil {
		slog.Error("[PathUser] failed to execute user update query", "error", errExe)
		if errCode := ur.db.ErrorCode(errExe); errCode == "1062" {
			return nil, domain.ErrConflictingData
		}
		return nil, errExe
	}

	rowsAffected, errRow := result.RowsAffected()
	if errRow != nil {
		slog.Error("[PathUser] failed to update user rows affected", "error", errRow)
		return nil, errRow
	}

	if rowsAffected == 0 {
		slog.Warn("[PathUser] user rows were not updated", "rows", rowsAffected)
		return user, nil
	}

	return ur.GetUserByID(ctx, user.ID)
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	query := ur.db.QueryBuilder.
		Delete("users").
		Where(sq.Eq{"usr_id": id})

	sqlQuery, args, errSql := query.ToSql()
	if errSql != nil {
		slog.Error("[DeleteUser] failed to generate sql query", "error", errSql)
		return errSql
	}

	result, err := ur.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[DeleteUser] failed to delete user", "id", id, "error", err)
		return err
	}

	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		slog.Error("[DeleteUser] failed to delete user rows affected", "error", errRows)
		return errRows
	}

	if rowsAffected == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}

func (ur *UserRepository) findUsers(ctx context.Context, query sq.SelectBuilder) ([]domain.User, error) {
	var users []domain.User

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[findUsers] failed to generate sql query", "error", err)
		return nil, err
	}

	rows, errQuery := ur.db.QueryContext(ctx, sqlQuery, args...)
	if errQuery != nil {
		slog.Error("[findUsers] failed to execute user query", "error", errQuery)
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User

		err := rows.Scan(
			&user.ID,
			&user.RoleID,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			slog.Error("[findUsers] failed to scan user row", "error", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[findUsers] failed to scan user rows", "error", err)
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) baseUserQuery() sq.SelectBuilder {
	return sq.Select(
		"usr_id",
		"rol_id",
		"usr_name",
		"usr_email",
		"usr_password_hash",
		"usr_created_at",
		"usr_updated_at",
	).From("users")
}
