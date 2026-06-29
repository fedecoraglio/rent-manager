package repository

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
)

type LocationRepository struct {
	db *mysql.DB
}

func NewLocationRepository(db *mysql.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}

func (lr *LocationRepository) ListCountries(ctx context.Context) ([]domain.Country, error) {
	query := lr.db.QueryBuilder.
		Select(
			"cou_id",
			"cou_code",
			"cou_name",
			"cou_created_at",
			"cou_updated_at",
		).
		From("countries").
		OrderBy("cou_name")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[ListCountries] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := lr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[ListCountries] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var countries []domain.Country

	for rows.Next() {
		var country domain.Country

		if err := rows.Scan(
			&country.ID,
			&country.Code,
			&country.Name,
			&country.CreatedAt,
			&country.UpdatedAt,
		); err != nil {
			slog.Error("[ListCountries] failed to scan row", "err", err)
			return nil, err
		}

		countries = append(countries, country)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[ListCountries] rows iteration failed", "err", err)
		return nil, err
	}

	return countries, nil
}

func (lr *LocationRepository) ListStatesByCountry(ctx context.Context, countryID int64) ([]domain.State, error) {
	query := lr.db.QueryBuilder.
		Select(
			"sta_id",
			"cou_id",
			"sta_code",
			"sta_name",
			"sta_created_at",
			"sta_updated_at",
		).
		From("states").
		Where("cou_id = ?", countryID).
		OrderBy("sta_name")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[ListStatesByCountry] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := lr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[ListStatesByCountry] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var states []domain.State

	for rows.Next() {
		var state domain.State

		if err := rows.Scan(
			&state.ID,
			&state.CountryID,
			&state.Code,
			&state.Name,
			&state.CreatedAt,
			&state.UpdatedAt,
		); err != nil {
			slog.Error("[ListStatesByCountry] failed to scan row", "err", err)
			return nil, err
		}

		states = append(states, state)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[ListStatesByCountry] rows iteration failed", "err", err)
		return nil, err
	}

	return states, nil
}
