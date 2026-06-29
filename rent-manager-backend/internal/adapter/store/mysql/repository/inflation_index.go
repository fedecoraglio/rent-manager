package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	sq "github.com/Masterminds/squirrel"

	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
)

type InflationIndexRepository struct {
	db *mysql.DB
}

func NewInflationIndexRepository(db *mysql.DB) *InflationIndexRepository {
	return &InflationIndexRepository{
		db: db,
	}
}

func (r *InflationIndexRepository) Create(
	ctx context.Context,
	inflationIndex *domain.InflationIndex,
) error {
	query := r.db.QueryBuilder.
		Insert("inflation_indexes").
		Columns(
			"ixi_period",
			"ixi_percentage",
			"ixi_source",
			"ixi_notes",
		).
		Values(
			inflationIndex.Period,
			inflationIndex.Percentage,
			inflationIndex.Source,
			inflationIndex.Notes,
		)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[CreateInflationIndex] failed to generate sql query", "err", err)
		return err
	}

	result, err := r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[CreateInflationIndex] failed to execute sql query", "err", err)

		if errCode := r.db.ErrorCode(err); errCode == "1062" {
			return domain.ErrInflationIndexAlreadyExists
		}

		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		slog.Error("[CreateInflationIndex] failed to get last insert id", "err", err)
		return err
	}

	created, err := r.GetByID(ctx, lastID)
	if err != nil {
		return err
	}

	*inflationIndex = *created

	return nil
}

func (r *InflationIndexRepository) Update(
	ctx context.Context,
	inflationIndex *domain.InflationIndex,
) error {
	query := r.db.QueryBuilder.
		Update("inflation_indexes").
		Set("ixi_period", inflationIndex.Period).
		Set("ixi_percentage", inflationIndex.Percentage).
		Set("ixi_source", inflationIndex.Source).
		Set("ixi_notes", inflationIndex.Notes).
		Where(sq.Eq{"ixi_id": inflationIndex.ID})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[UpdateInflationIndex] failed to generate sql query", "err", err)
		return err
	}

	result, err := r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[UpdateInflationIndex] failed to execute sql query", "err", err)

		if errCode := r.db.ErrorCode(err); errCode == "1062" {
			return domain.ErrInflationIndexAlreadyExists
		}

		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("[UpdateInflationIndex] failed to get rows affected", "err", err)
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrInflationIndexNotFound
	}

	updated, err := r.GetByID(ctx, inflationIndex.ID)
	if err != nil {
		return err
	}

	*inflationIndex = *updated

	return nil
}

func (r *InflationIndexRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.InflationIndex, error) {
	query := r.baseInflationIndexQuery().
		Where(sq.Eq{"ixi_id": id}).
		Limit(1)

	return r.findOne(ctx, query)
}

func (r *InflationIndexRepository) GetByPeriod(
	ctx context.Context,
	period time.Time,
) (*domain.InflationIndex, error) {
	query := r.baseInflationIndexQuery().
		Where(sq.Eq{"ixi_period": period}).
		Limit(1)

	return r.findOne(ctx, query)
}

func (r *InflationIndexRepository) List(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.InflationIndex, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := r.baseInflationIndexQuery().
		OrderBy("ixi_period DESC").
		Limit(limit).
		Offset(offset)

	return r.findMany(ctx, query)
}

func (r *InflationIndexRepository) ListByPeriodRange(
	ctx context.Context,
	from time.Time,
	to time.Time,
) ([]domain.InflationIndex, error) {
	query := r.baseInflationIndexQuery().
		Where(sq.GtOrEq{"ixi_period": from}).
		Where(sq.LtOrEq{"ixi_period": to}).
		OrderBy("ixi_period ASC")

	return r.findMany(ctx, query)
}

func (r *InflationIndexRepository) ListLatestBeforePeriod(
	ctx context.Context,
	period time.Time,
	limit uint64,
) ([]domain.InflationIndex, error) {
	query := r.baseInflationIndexQuery().
		Where(sq.Lt{"ixi_period": period}).
		OrderBy("ixi_period DESC").
		Limit(limit)

	inflationIndexes, err := r.findMany(ctx, query)
	if err != nil {
		return nil, err
	}

	// Reverse the slice to keep the indexes in chronological order.
	for i, j := 0, len(inflationIndexes)-1; i < j; i, j = i+1, j-1 {
		inflationIndexes[i], inflationIndexes[j] = inflationIndexes[j], inflationIndexes[i]
	}

	return inflationIndexes, nil
}

func (r *InflationIndexRepository) baseInflationIndexQuery() sq.SelectBuilder {
	return r.db.QueryBuilder.
		Select(
			"ixi_id",
			"ixi_period",
			"ixi_percentage",
			"ixi_source",
			"ixi_notes",
			"ixi_created_at",
			"ixi_updated_at",
		).
		From("inflation_indexes")
}

func (r *InflationIndexRepository) findOne(
	ctx context.Context,
	query sq.SelectBuilder,
) (*domain.InflationIndex, error) {
	inflationIndexes, err := r.findMany(ctx, query)
	if err != nil {
		return nil, err
	}

	if len(inflationIndexes) == 0 {
		return nil, domain.ErrInflationIndexNotFound
	}

	return &inflationIndexes[0], nil
}

func (r *InflationIndexRepository) findMany(
	ctx context.Context,
	query sq.SelectBuilder,
) ([]domain.InflationIndex, error) {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[InflationIndexRepository] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[InflationIndexRepository] failed to execute sql query", "err", err)

		if errors.Is(err, sql.ErrNoRows) {
			return []domain.InflationIndex{}, nil
		}

		return nil, err
	}
	defer rows.Close()

	inflationIndexes := make([]domain.InflationIndex, 0)

	for rows.Next() {
		var inflationIndex domain.InflationIndex

		if err := rows.Scan(
			&inflationIndex.ID,
			&inflationIndex.Period,
			&inflationIndex.Percentage,
			&inflationIndex.Source,
			&inflationIndex.Notes,
			&inflationIndex.CreatedAt,
			&inflationIndex.UpdatedAt,
		); err != nil {
			slog.Error("[InflationIndexRepository] failed to scan row", "err", err)
			return nil, err
		}

		inflationIndexes = append(inflationIndexes, inflationIndex)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[InflationIndexRepository] rows error", "err", err)
		return nil, err
	}

	return inflationIndexes, nil
}
