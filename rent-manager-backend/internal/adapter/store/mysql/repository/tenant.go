package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"rent-manager-backend/internal/adapter/store/mysql/repository/shared"

	sq "github.com/Masterminds/squirrel"

	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
)

type TenantRepository struct {
	db *mysql.DB
}

func NewTenantRepository(db *mysql.DB) *TenantRepository {
	return &TenantRepository{
		db: db,
	}
}

func (tr *TenantRepository) CreateTenant(ctx context.Context, tenant *domain.Tenant) (*domain.Tenant, error) {
	query := tr.db.QueryBuilder.
		Insert("tenants").
		Columns(
			"cou_id",
			"sta_id",
			"ten_name",
			"ten_email",
			"ten_phone",
			"ten_document_number",
			"ten_city",
			"ten_street",
			"ten_street_number",
			"ten_floor",
			"ten_apartment",
			"ten_postal_code",
		).
		Values(
			shared.NullInt64(tenant.CountryID),
			shared.NullInt64(tenant.StateID),
			tenant.Name,
			tenant.Email,
			tenant.Phone,
			tenant.DocumentNumber,
			tenant.City,
			tenant.Street,
			tenant.StreetNumber,
			tenant.Floor,
			tenant.Apartment,
			tenant.PostalCode,
		)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[CreateTenant] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := tr.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[CreateTenant] failed to execute sql query", "err", err)

		if errCode := tr.db.ErrorCode(err); errCode == "1062" {
			return nil, domain.ErrTenantAlreadyExists
		}

		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		slog.Error("[CreateTenant] failed to get last insert id", "err", err)
		return nil, err
	}

	return tr.GetTenantByID(ctx, lastID)
}

func (tr *TenantRepository) GetTenantByID(ctx context.Context, id int64) (*domain.Tenant, error) {
	var tenant domain.Tenant

	query := tr.baseTenantQuery().
		Where(sq.Eq{"ten_id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetTenantByID] failed to generate sql query", "err", err)
		return nil, err
	}

	err = tr.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&tenant.ID,
		&tenant.CountryID,
		&tenant.StateID,
		&tenant.Name,
		&tenant.Email,
		&tenant.Phone,
		&tenant.DocumentNumber,
		&tenant.City,
		&tenant.Street,
		&tenant.StreetNumber,
		&tenant.Floor,
		&tenant.Apartment,
		&tenant.PostalCode,
		&tenant.CreatedAt,
		&tenant.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetTenantByID] failed to scan row", "err", err)
		return nil, err
	}

	return &tenant, nil
}

func (tr *TenantRepository) GetTenantByDocumentNumber(ctx context.Context, docNumber string) (*domain.Tenant, error) {
	var tenant domain.Tenant

	query := tr.baseTenantQuery().
		Where(sq.Eq{"ten_document_number": docNumber}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetTenantByDocumentNumber] failed to generate sql query", "err", err)
		return nil, err
	}

	err = tr.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&tenant.ID,
		&tenant.CountryID,
		&tenant.StateID,
		&tenant.Name,
		&tenant.Email,
		&tenant.Phone,
		&tenant.DocumentNumber,
		&tenant.City,
		&tenant.Street,
		&tenant.StreetNumber,
		&tenant.Floor,
		&tenant.Apartment,
		&tenant.PostalCode,
		&tenant.CreatedAt,
		&tenant.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetTenantByDocumentNumber] failed to scan row", "err", err)
		return nil, err
	}

	return &tenant, nil
}

func (tr *TenantRepository) ListTenants(ctx context.Context, page uint64, limit uint64) ([]domain.Tenant, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := tr.baseTenantQuery().
		OrderBy("ten_id").
		Limit(limit).
		Offset(offset)

	return tr.findTenants(ctx, query)
}

func (tr *TenantRepository) SearchTenants(
	ctx context.Context,
	value string,
	page uint64,
	limit uint64,
) ([]domain.Tenant, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	searchValue := "%" + value + "%"

	query := tr.baseTenantQuery().
		Where(
			sq.Or{
				sq.Expr("LOWER(ten_name) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(ten_email) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(ten_phone) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(ten_document_number) LIKE LOWER(?)", searchValue),
				sq.Expr("LOWER(ten_city) LIKE LOWER(?)", searchValue),
			},
		).
		OrderBy("ten_id").
		Limit(limit).
		Offset(offset)

	return tr.findTenants(ctx, query)
}

func (tr *TenantRepository) UpdateTenant(
	ctx context.Context,
	tenant *domain.Tenant,
) (*domain.Tenant, error) {
	queryBuilder := tr.db.QueryBuilder.Update("tenants")

	if tenant.CountryID != 0 {
		queryBuilder = queryBuilder.Set("cou_id", tenant.CountryID)
	}

	if tenant.StateID != 0 {
		queryBuilder = queryBuilder.Set("sta_id", tenant.StateID)
	}

	if tenant.Name != "" {
		queryBuilder = queryBuilder.Set("ten_name", tenant.Name)
	}

	if tenant.Email != "" {
		queryBuilder = queryBuilder.Set("ten_email", tenant.Email)
	}

	if tenant.Phone != "" {
		queryBuilder = queryBuilder.Set("ten_phone", tenant.Phone)
	}

	if tenant.DocumentNumber != "" {
		queryBuilder = queryBuilder.Set("ten_document_number", tenant.DocumentNumber)
	}

	if tenant.City != "" {
		queryBuilder = queryBuilder.Set("ten_city", tenant.City)
	}

	if tenant.Street != "" {
		queryBuilder = queryBuilder.Set("ten_street", tenant.Street)
	}

	if tenant.StreetNumber != "" {
		queryBuilder = queryBuilder.Set("ten_street_number", tenant.StreetNumber)
	}

	if tenant.Floor != "" {
		queryBuilder = queryBuilder.Set("ten_floor", tenant.Floor)
	}

	if tenant.Apartment != "" {
		queryBuilder = queryBuilder.Set("ten_apartment", tenant.Apartment)
	}

	if tenant.PostalCode != "" {
		queryBuilder = queryBuilder.Set("ten_postal_code", tenant.PostalCode)
	}

	queryBuilder = queryBuilder.Where(sq.Eq{"ten_id": tenant.ID})

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		slog.Error("[UpdateTenant] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := tr.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[UpdateTenant] failed to execute sql query", "err", err)

		if errCode := tr.db.ErrorCode(err); errCode == "1062" {
			return nil, domain.ErrTenantAlreadyExists
		}

		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("[UpdateTenant] failed to get rows affected", "err", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domain.ErrDataNotFound
	}

	return tr.GetTenantByID(ctx, tenant.ID)
}

func (tr *TenantRepository) findTenants(ctx context.Context, query sq.SelectBuilder) ([]domain.Tenant, error) {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[findTenants] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := tr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[findTenants] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var tenants []domain.Tenant

	for rows.Next() {
		var tenant domain.Tenant

		err := rows.Scan(
			&tenant.ID,
			&tenant.CountryID,
			&tenant.StateID,
			&tenant.Name,
			&tenant.Email,
			&tenant.Phone,
			&tenant.DocumentNumber,
			&tenant.City,
			&tenant.Street,
			&tenant.StreetNumber,
			&tenant.Floor,
			&tenant.Apartment,
			&tenant.PostalCode,
			&tenant.CreatedAt,
			&tenant.UpdatedAt,
		)
		if err != nil {
			slog.Error("[findTenants] failed to scan row", "err", err)
			return nil, err
		}

		tenants = append(tenants, tenant)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[findTenants] rows iteration failed", "err", err)
		return nil, err
	}

	return tenants, nil
}

func (tr *TenantRepository) baseTenantQuery() sq.SelectBuilder {
	return tr.db.QueryBuilder.
		Select(
			"ten_id",
			"COALESCE(cou_id, 0)",
			"COALESCE(sta_id, 0)",
			"ten_name",
			"COALESCE(ten_email, '')",
			"COALESCE(ten_phone, '')",
			"ten_document_number",
			"COALESCE(ten_city, '')",
			"COALESCE(ten_street, '')",
			"COALESCE(ten_street_number, '')",
			"COALESCE(ten_floor, '')",
			"COALESCE(ten_apartment, '')",
			"COALESCE(ten_postal_code, '')",
			"ten_created_at",
			"ten_updated_at",
		).
		From("tenants")
}
