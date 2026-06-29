package repository

import (
	"context"
	"log/slog"

	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
)

type ContractCatalogRepository struct {
	db *mysql.DB
}

func NewContractCatalogRepository(db *mysql.DB) *ContractCatalogRepository {
	return &ContractCatalogRepository{
		db: db,
	}
}

func (ccr *ContractCatalogRepository) ListContractStatuses(
	ctx context.Context,
) ([]domain.ContractStatus, error) {
	query := ccr.db.QueryBuilder.
		Select(
			"cst_id",
			"cst_code",
			"cst_name",
		).
		From("contract_statuses").
		OrderBy("cst_name")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[ListContractStatuses] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := ccr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[ListContractStatuses] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var statuses []domain.ContractStatus

	for rows.Next() {
		var status domain.ContractStatus

		if err := rows.Scan(
			&status.ID,
			&status.Code,
			&status.Name,
		); err != nil {
			slog.Error("[ListContractStatuses] failed to scan row", "err", err)
			return nil, err
		}

		statuses = append(statuses, status)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[ListContractStatuses] rows iteration failed", "err", err)
		return nil, err
	}

	return statuses, nil
}

func (ccr *ContractCatalogRepository) ListInterestCalculationTypes(
	ctx context.Context,
) ([]domain.InterestCalculationType, error) {
	query := ccr.db.QueryBuilder.
		Select(
			"ict_id",
			"ict_code",
			"ict_name",
		).
		From("interest_calculation_types").
		OrderBy("ict_name")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[ListInterestCalculationTypes] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := ccr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[ListInterestCalculationTypes] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var types []domain.InterestCalculationType

	for rows.Next() {
		var interestType domain.InterestCalculationType

		if err := rows.Scan(
			&interestType.ID,
			&interestType.Code,
			&interestType.Name,
		); err != nil {
			slog.Error("[ListInterestCalculationTypes] failed to scan row", "err", err)
			return nil, err
		}

		types = append(types, interestType)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[ListInterestCalculationTypes] rows iteration failed", "err", err)
		return nil, err
	}

	return types, nil
}

func (ccr *ContractCatalogRepository) ListRentAdjustmentTypes(
	ctx context.Context,
) ([]domain.RentAdjustmentType, error) {
	query := ccr.db.QueryBuilder.
		Select(
			"rat_id",
			"rat_code",
			"rat_name",
		).
		From("rent_adjustment_types").
		OrderBy("rat_name")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[ListRentAdjustmentTypes] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := ccr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[ListRentAdjustmentTypes] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var types []domain.RentAdjustmentType

	for rows.Next() {
		var adjustmentType domain.RentAdjustmentType

		if err := rows.Scan(
			&adjustmentType.ID,
			&adjustmentType.Code,
			&adjustmentType.Name,
		); err != nil {
			slog.Error("[ListRentAdjustmentTypes] failed to scan row", "err", err)
			return nil, err
		}

		types = append(types, adjustmentType)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[ListRentAdjustmentTypes] rows iteration failed", "err", err)
		return nil, err
	}

	return types, nil
}
