package repository

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
)

type PropertyCatalogRepository struct {
	db *mysql.DB
}

func NewPropertyCatalogRepository(db *mysql.DB) *PropertyCatalogRepository {
	return &PropertyCatalogRepository{
		db: db,
	}
}

func (pr *PropertyCatalogRepository) ListPropertyTypes(ctx context.Context) ([]domain.PropertyType, error) {
	query := pr.db.QueryBuilder.
		Select(
			"pty_id",
			"pty_code",
			"pty_name",
		).
		From("property_types").
		OrderBy("pty_name")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[ListPropertyTypes] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := pr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[ListPropertyTypes] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var propertyTypes []domain.PropertyType

	for rows.Next() {
		var propertyType domain.PropertyType

		if err := rows.Scan(
			&propertyType.ID,
			&propertyType.Code,
			&propertyType.Name,
		); err != nil {
			slog.Error("[ListPropertyTypes] failed to scan row", "err", err)
			return nil, err
		}

		propertyTypes = append(propertyTypes, propertyType)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[ListPropertyTypes] rows iteration failed", "err", err)
		return nil, err
	}

	return propertyTypes, nil
}

func (pr *PropertyCatalogRepository) ListPropertyStatuses(ctx context.Context) ([]domain.PropertyStatus, error) {
	query := pr.db.QueryBuilder.
		Select(
			"pst_id",
			"pst_code",
			"pst_name",
		).
		From("property_statuses").
		OrderBy("pst_name")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[ListPropertyStatuses] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := pr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[ListPropertyStatuses] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var propertyStatuses []domain.PropertyStatus

	for rows.Next() {
		var propertyStatus domain.PropertyStatus

		if err := rows.Scan(
			&propertyStatus.ID,
			&propertyStatus.Code,
			&propertyStatus.Name,
		); err != nil {
			slog.Error("[ListPropertyStatuses] failed to scan row", "err", err)
			return nil, err
		}

		propertyStatuses = append(propertyStatuses, propertyStatus)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[ListPropertyStatuses] rows iteration failed", "err", err)
		return nil, err
	}

	return propertyStatuses, nil
}
