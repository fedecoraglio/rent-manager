package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	sq "github.com/Masterminds/squirrel"

	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
)

type OwnerRepository struct {
	db *mysql.DB
}

func NewOwnerRepository(db *mysql.DB) *OwnerRepository {
	return &OwnerRepository{
		db: db,
	}
}

func (or *OwnerRepository) CreateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
	query := or.db.QueryBuilder.
		Insert("owners").
		Columns(
			"own_name",
			"own_email",
			"own_phone",
			"own_document_number",
		).
		Values(
			owner.Name,
			owner.Email,
			owner.Phone,
			owner.DocumentNumber,
		)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[CreateOwner] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := or.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[CreateOwner] failed to execute sql query", "err", err)

		if errCode := or.db.ErrorCode(err); errCode == "1062" {
			return nil, domain.ErrOwnerAlreadyExists
		}

		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		slog.Error("[CreateOwner] failed to get last insert id", "err", err)
		return nil, err
	}

	return or.GetOwnerByID(ctx, lastID)
}

func (or *OwnerRepository) GetOwnerByID(ctx context.Context, id int64) (*domain.Owner, error) {
	var owner domain.Owner

	query := or.baseOwnerQuery().
		Where(sq.Eq{"own_id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetOwnerByID] failed to generate sql query", "err", err)
		return nil, err
	}

	err = or.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&owner.ID,
		&owner.Name,
		&owner.Email,
		&owner.Phone,
		&owner.DocumentNumber,
		&owner.CreatedAt,
		&owner.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetOwnerByID] failed to scan row", "err", err)
		return nil, err
	}

	return &owner, nil
}

func (or *OwnerRepository) GetOwnerByDocumentNumber(ctx context.Context, documentNumber string) (*domain.Owner, error) {
	var owner domain.Owner

	query := or.baseOwnerQuery().
		Where(sq.Eq{"own_document_number": documentNumber}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetOwnerByDocumentNumber] failed to generate sql query", "err", err)
		return nil, err
	}

	err = or.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&owner.ID,
		&owner.Name,
		&owner.Email,
		&owner.Phone,
		&owner.DocumentNumber,
		&owner.CreatedAt,
		&owner.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetOwnerByDocumentNumber] failed to scan row", "err", err)
		return nil, err
	}

	return &owner, nil
}

func (or *OwnerRepository) ListOwners(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.Owner, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := or.baseOwnerQuery().
		OrderBy("own_id").
		Limit(limit).
		Offset(offset)

	return or.findOwners(ctx, query)
}

func (or *OwnerRepository) SearchOwners(
	ctx context.Context,
	value string,
	page uint64,
	limit uint64,
) ([]domain.Owner, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	searchValue := "%" + value + "%"

	query := or.baseOwnerQuery().
		Where(
			sq.Or{
				sq.Expr("LOWER(own_name) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(own_email) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(own_phone) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(own_document_number) LIKE LOWER(?)", searchValue),
			},
		).
		OrderBy("own_id").
		Limit(limit).
		Offset(offset)

	return or.findOwners(ctx, query)
}

func (or *OwnerRepository) UpdateOwner(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
	queryBuilder := or.db.QueryBuilder.Update("owners")

	if owner.Name != "" {
		queryBuilder = queryBuilder.Set("own_name", owner.Name)
	}

	if owner.Email != "" {
		queryBuilder = queryBuilder.Set("own_email", owner.Email)
	}

	if owner.Phone != "" {
		queryBuilder = queryBuilder.Set("own_phone", owner.Phone)
	}

	if owner.DocumentNumber != "" {
		queryBuilder = queryBuilder.Set("own_document_number", owner.DocumentNumber)
	}

	queryBuilder = queryBuilder.Where(sq.Eq{"own_id": owner.ID})

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		slog.Error("[UpdateOwner] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := or.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[UpdateOwner] failed to execute sql query", "err", err)

		if errCode := or.db.ErrorCode(err); errCode == "1062" {
			return nil, domain.ErrOwnerAlreadyExists
		}

		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("[UpdateOwner] failed to get rows affected", "err", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domain.ErrDataNotFound
	}

	return or.GetOwnerByID(ctx, owner.ID)
}

func (or *OwnerRepository) baseOwnerQuery() sq.SelectBuilder {
	return or.db.QueryBuilder.
		Select(
			"own_id",
			"own_name",
			"own_email",
			"own_phone",
			"own_document_number",
			"own_created_at",
			"own_updated_at",
		).
		From("owners")
}

func (or *OwnerRepository) findOwners(ctx context.Context, query sq.SelectBuilder) ([]domain.Owner, error) {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[findOwners] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := or.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[findOwners] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var owners []domain.Owner

	for rows.Next() {
		var owner domain.Owner

		err := rows.Scan(
			&owner.ID,
			&owner.Name,
			&owner.Email,
			&owner.Phone,
			&owner.DocumentNumber,
			&owner.CreatedAt,
			&owner.UpdatedAt,
		)
		if err != nil {
			slog.Error("[findOwners] failed to scan row", "err", err)
			return nil, err
		}

		owners = append(owners, owner)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[findOwners] rows iteration failed", "err", err)
		return nil, err
	}

	return owners, nil
}
