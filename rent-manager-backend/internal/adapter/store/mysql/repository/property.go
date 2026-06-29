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

type PropertyRepository struct {
	db *mysql.DB
}

func NewPropertyRepository(db *mysql.DB) *PropertyRepository {
	return &PropertyRepository{
		db: db,
	}
}

func (pr *PropertyRepository) CreateProperty(ctx context.Context, property *domain.Property) (*domain.Property, error) {
	query := pr.db.QueryBuilder.
		Insert("properties").
		Columns(
			"own_id",
			"pty_id",
			"pst_id",
			"cou_id",
			"sta_id",
			"pro_code",
			"pro_title",
			"pro_description",
			"pro_street",
			"pro_street_number",
			"pro_floor",
			"pro_apartment",
			"pro_city",
			"pro_postal_code",
		).
		Values(
			property.OwnerID,
			property.TypeID,
			property.StatusID,
			property.CountryID,
			property.StateID,
			property.Code,
			property.Title,
			property.Description,
			property.Street,
			property.StreetNumber,
			property.Floor,
			property.Apartment,
			property.City,
			property.PostalCode,
		)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[CreateProperty] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := pr.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[CreateProperty] failed to execute sql query", "err", err)

		if errCode := pr.db.ErrorCode(err); errCode == "1062" {
			return nil, domain.ErrPropertyAlreadyExists
		}

		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		slog.Error("[CreateProperty] failed to get last insert id", "err", err)
		return nil, err
	}

	return pr.GetPropertyByID(ctx, lastID)
}

func (pr *PropertyRepository) GetPropertyByID(ctx context.Context, id int64) (*domain.Property, error) {
	var property domain.Property

	query := pr.basePropertyQuery().
		Where(sq.Eq{"p.pro_id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetPropertyByID] failed to generate sql query", "err", err)
		return nil, err
	}

	err = pr.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&property.ID,
		&property.OwnerID,
		&property.TypeID,
		&property.StatusID,
		&property.CountryID,
		&property.StateID,
		&property.Code,
		&property.Title,
		&property.Description,
		&property.Street,
		&property.StreetNumber,
		&property.Floor,
		&property.Apartment,
		&property.City,
		&property.PostalCode,
		&property.CreatedAt,
		&property.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetPropertyByID] failed to scan row", "err", err)
		return nil, err
	}

	return &property, nil
}

func (pr *PropertyRepository) GetPropertyByCode(ctx context.Context, code string) (*domain.Property, error) {
	var property domain.Property

	query := pr.basePropertyQuery().
		Where(sq.Eq{"p.pro_code": code}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetPropertyByCode] failed to generate sql query", "err", err)
		return nil, err
	}

	err = pr.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&property.ID,
		&property.OwnerID,
		&property.TypeID,
		&property.StatusID,
		&property.CountryID,
		&property.StateID,
		&property.Code,
		&property.Title,
		&property.Description,
		&property.Street,
		&property.StreetNumber,
		&property.Floor,
		&property.Apartment,
		&property.City,
		&property.PostalCode,
		&property.CreatedAt,
		&property.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetPropertyByCode] failed to scan row", "err", err)
		return nil, err
	}

	return &property, nil
}

func (pr *PropertyRepository) ListProperties(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.Property, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := pr.basePropertyQuery().
		OrderBy("p.pro_id").
		Limit(limit).
		Offset(offset)

	return pr.findProperties(ctx, query)
}

func (pr *PropertyRepository) SearchProperties(
	ctx context.Context,
	value string,
	page uint64,
	limit uint64,
) ([]domain.Property, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	searchValue := "%" + value + "%"

	query := pr.basePropertyQuery().
		Where(
			sq.Or{
				sq.Expr("LOWER(p.pro_code) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(p.pro_title) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(p.pro_description) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(p.pro_street) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(p.pro_city) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(p.pro_postal_code) LIKE LOWER(?)", searchValue),
			},
		).
		OrderBy("p.pro_id").
		Limit(limit).
		Offset(offset)

	return pr.findProperties(ctx, query)
}

func (pr *PropertyRepository) UpdateProperty(
	ctx context.Context,
	property *domain.Property,
) (*domain.Property, error) {
	queryBuilder := pr.db.QueryBuilder.Update("properties")

	if property.OwnerID != 0 {
		queryBuilder = queryBuilder.Set("own_id", property.OwnerID)
	}

	if property.TypeID != 0 {
		queryBuilder = queryBuilder.Set("pty_id", property.TypeID)
	}

	if property.StatusID != 0 {
		queryBuilder = queryBuilder.Set("pst_id", property.StatusID)
	}

	if property.CountryID != 0 {
		queryBuilder = queryBuilder.Set("cou_id", property.CountryID)
	}

	if property.StateID != 0 {
		queryBuilder = queryBuilder.Set("sta_id", property.StateID)
	}

	if property.Code != "" {
		queryBuilder = queryBuilder.Set("pro_code", property.Code)
	}

	if property.Title != "" {
		queryBuilder = queryBuilder.Set("pro_title", property.Title)
	}

	if property.Description != "" {
		queryBuilder = queryBuilder.Set("pro_description", property.Description)
	}

	if property.Street != "" {
		queryBuilder = queryBuilder.Set("pro_street", property.Street)
	}

	if property.StreetNumber != "" {
		queryBuilder = queryBuilder.Set("pro_street_number", property.StreetNumber)
	}

	if property.Floor != "" {
		queryBuilder = queryBuilder.Set("pro_floor", property.Floor)
	}

	if property.Apartment != "" {
		queryBuilder = queryBuilder.Set("pro_apartment", property.Apartment)
	}

	if property.City != "" {
		queryBuilder = queryBuilder.Set("pro_city", property.City)
	}

	if property.PostalCode != "" {
		queryBuilder = queryBuilder.Set("pro_postal_code", property.PostalCode)
	}

	queryBuilder = queryBuilder.Where(sq.Eq{"pro_id": property.ID})

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		slog.Error("[UpdateProperty] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := pr.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[UpdateProperty] failed to execute sql query", "err", err)

		if errCode := pr.db.ErrorCode(err); errCode == "1062" {
			return nil, domain.ErrPropertyAlreadyExists
		}

		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("[UpdateProperty] failed to get rows affected", "err", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domain.ErrDataNotFound
	}

	return pr.GetPropertyByID(ctx, property.ID)
}

func (pr *PropertyRepository) basePropertyQuery() sq.SelectBuilder {
	return pr.db.QueryBuilder.
		Select(
			"p.pro_id",
			"p.own_id",
			"p.pty_id",
			"p.pst_id",
			"p.cou_id",
			"p.sta_id",
			"p.pro_code",
			"p.pro_title",
			"COALESCE(p.pro_description, '')",
			"p.pro_street",
			"COALESCE(p.pro_street_number, '')",
			"COALESCE(p.pro_floor, '')",
			"COALESCE(p.pro_apartment, '')",
			"p.pro_city",
			"COALESCE(p.pro_postal_code, '')",
			"p.pro_created_at",
			"p.pro_updated_at",
		).
		From("properties p")
}

func (pr *PropertyRepository) findProperties(ctx context.Context, query sq.SelectBuilder) ([]domain.Property, error) {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[findProperties] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := pr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[findProperties] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var properties []domain.Property

	for rows.Next() {
		var property domain.Property

		err := rows.Scan(
			&property.ID,
			&property.OwnerID,
			&property.TypeID,
			&property.StatusID,
			&property.CountryID,
			&property.StateID,
			&property.Code,
			&property.Title,
			&property.Description,
			&property.Street,
			&property.StreetNumber,
			&property.Floor,
			&property.Apartment,
			&property.City,
			&property.PostalCode,
			&property.CreatedAt,
			&property.UpdatedAt,
		)
		if err != nil {
			slog.Error("[findProperties] failed to scan row", "err", err)
			return nil, err
		}

		properties = append(properties, property)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[findProperties] rows iteration failed", "err", err)
		return nil, err
	}

	return properties, nil
}
