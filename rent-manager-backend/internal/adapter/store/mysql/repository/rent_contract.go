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

type RentalContractRepository struct {
	db *mysql.DB
}

func NewRentalContractRepository(db *mysql.DB) *RentalContractRepository {
	return &RentalContractRepository{db: db}
}

func (rcr *RentalContractRepository) CreateRentalContract(
	ctx context.Context,
	rentalContract *domain.RentalContract,
) (*domain.RentalContract, error) {
	query := rcr.db.QueryBuilder.
		Insert("rental_contracts").
		Columns(
			"pro_id",
			"ten_id",
			"cst_id",
			"ict_id",
			"rat_id",
			"rco_start_date",
			"rco_end_date",
			"rco_total_payments",
			"rco_monthly_amount",
			"rco_deposit_amount",
			"rco_currency",
			"rco_due_day",
			"rco_daily_interest_percentage",
			"rco_adjustment_frequency_months",
			"rco_notes",
		).
		Values(
			rentalContract.PropertyID,
			rentalContract.TenantID,
			rentalContract.StatusID,
			rentalContract.InterestCalculationTypeID,
			rentalContract.AdjustmentTypeID,
			rentalContract.StartDate,
			rentalContract.EndDate,
			rentalContract.TotalPayments,
			rentalContract.MonthlyAmount,
			rentalContract.DepositAmount,
			rentalContract.Currency,
			rentalContract.DueDay,
			rentalContract.DailyInterestPercentage,
			rentalContract.AdjustmentFrequencyMonths,
			rentalContract.Notes,
		)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[CreateRentalContract] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := rcr.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[CreateRentalContract] failed to execute sql query", "err", err)
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		slog.Error("[CreateRentalContract] failed to get last insert id", "err", err)
		return nil, err
	}

	return rcr.GetRentalContractByID(ctx, lastID)
}

func (rcr *RentalContractRepository) GetRentalContractByID(
	ctx context.Context,
	id int64,
) (*domain.RentalContract, error) {
	var rentalContract domain.RentalContract

	query := rcr.baseRentalContractQuery().
		Where(sq.Eq{"rc.rco_id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetRentalContractByID] failed to generate sql query", "err", err)
		return nil, err
	}

	rentalContract, err = scanRentalContract(
		rcr.db.QueryRowContext(ctx, sqlQuery, args...),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetRentalContractByID] failed to scan row", "err", err)
		return nil, err
	}

	return &rentalContract, nil
}

func (rcr *RentalContractRepository) UpdateRentalContract(
	ctx context.Context,
	rentalContract *domain.RentalContract,
) (*domain.RentalContract, error) {
	queryBuilder := rcr.db.QueryBuilder.Update("rental_contracts")

	if rentalContract.PropertyID != 0 {
		queryBuilder = queryBuilder.Set("pro_id", rentalContract.PropertyID)
	}

	if rentalContract.TenantID != 0 {
		queryBuilder = queryBuilder.Set("ten_id", rentalContract.TenantID)
	}

	if rentalContract.StatusID != 0 {
		queryBuilder = queryBuilder.Set("cst_id", rentalContract.StatusID)
	}

	if rentalContract.InterestCalculationTypeID != 0 {
		queryBuilder = queryBuilder.Set("ict_id", rentalContract.InterestCalculationTypeID)
	}

	if rentalContract.AdjustmentTypeID != 0 {
		queryBuilder = queryBuilder.Set("rat_id", rentalContract.AdjustmentTypeID)
	}

	if !rentalContract.StartDate.IsZero() {
		queryBuilder = queryBuilder.Set("rco_start_date", rentalContract.StartDate)
	}

	if !rentalContract.EndDate.IsZero() {
		queryBuilder = queryBuilder.Set("rco_end_date", rentalContract.EndDate)
	}

	if rentalContract.MonthlyAmount != 0 {
		queryBuilder = queryBuilder.Set("rco_monthly_amount", rentalContract.MonthlyAmount)
	}

	if rentalContract.DepositAmount >= 0 {
		queryBuilder = queryBuilder.Set("rco_deposit_amount", rentalContract.DepositAmount)
	}

	if rentalContract.Currency != "" {
		queryBuilder = queryBuilder.Set("rco_currency", rentalContract.Currency)
	}

	if rentalContract.DueDay != 0 {
		queryBuilder = queryBuilder.Set("rco_due_day", rentalContract.DueDay)
	}

	if rentalContract.TotalPayments > 0 {
		queryBuilder = queryBuilder.Set("rco_total_payments", rentalContract.TotalPayments)
	}

	if rentalContract.DailyInterestPercentage >= 0 {
		queryBuilder = queryBuilder.Set(
			"rco_daily_interest_percentage",
			rentalContract.DailyInterestPercentage,
		)
	}

	if rentalContract.AdjustmentFrequencyMonths != 0 {
		queryBuilder = queryBuilder.Set(
			"rco_adjustment_frequency_months",
			rentalContract.AdjustmentFrequencyMonths,
		)
	}

	queryBuilder = queryBuilder.Set("rco_notes", rentalContract.Notes)
	queryBuilder = queryBuilder.Where(sq.Eq{"rco_id": rentalContract.ID})

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		slog.Error("[UpdateRentalContract] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := rcr.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[UpdateRentalContract] failed to execute sql query", "err", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("[UpdateRentalContract] failed to get rows affected", "err", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domain.ErrDataNotFound
	}

	return rcr.GetRentalContractByID(ctx, rentalContract.ID)
}

func (rcr *RentalContractRepository) ListRentalContracts(
	ctx context.Context,
	propertyID int64,
	page uint64,
	limit uint64,
) ([]domain.RentalContract, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := rcr.baseRentalContractQuery().
		OrderBy("rc.rco_id").
		Limit(limit).
		Offset(offset)

	if propertyID > 0 {
		query = query.Where(sq.Eq{"rc.pro_id": propertyID})
	}

	return rcr.findRentalContracts(ctx, query)
}

func (rcr *RentalContractRepository) GetActiveRentalContractByPropertyID(
	ctx context.Context,
	propertyID int64,
) (*domain.RentalContract, error) {
	var rentalContract domain.RentalContract

	query := rcr.baseRentalContractQuery().
		Join("contract_statuses cs ON cs.cst_id = rc.cst_id").
		Where(sq.Eq{"rc.pro_id": propertyID}).
		Where(sq.Eq{"cs.cst_code": "active"}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetActiveRentalContractByPropertyID] failed to generate sql query", "err", err)
		return nil, err
	}

	rentalContract, err = scanRentalContract(
		rcr.db.QueryRowContext(ctx, sqlQuery, args...),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetActiveRentalContractByPropertyID] failed to scan row", "err", err)
		return nil, err
	}

	return &rentalContract, nil
}

func (rcr *RentalContractRepository) ListActiveRentalContracts(
	ctx context.Context,
	page uint64,
	limit uint64,
) ([]domain.RentalContract, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	query := rcr.baseRentalContractQuery().
		Where(sq.Eq{"cst.cst_code": "active"}).
		Offset(offset).
		Limit(limit).
		OrderBy("p.pro_title")

	rentContracts, err := rcr.findRentalContracts(ctx, query)
	if err != nil {
		return nil, err
	}

	return rentContracts, nil
}

func (rcr *RentalContractRepository) findRentalContracts(
	ctx context.Context,
	query sq.SelectBuilder,
) ([]domain.RentalContract, error) {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[findRentalContracts] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := rcr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[findRentalContracts] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var rentalContracts []domain.RentalContract

	for rows.Next() {
		rentalContract, err := scanRentalContract(rows)
		if err != nil {
			slog.Error("[findRentalContracts] failed to scan row", "err", err)
			return nil, err
		}

		rentalContracts = append(rentalContracts, rentalContract)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[findRentalContracts] rows iteration failed", "err", err)
		return nil, err
	}

	return rentalContracts, nil
}

func (rcr *RentalContractRepository) baseRentalContractQuery() sq.SelectBuilder {
	return rcr.db.QueryBuilder.
		Select(
			"rc.rco_id",
			"rc.pro_id",
			"rc.ten_id",
			"rc.cst_id",
			"rc.ict_id",
			"rc.rat_id",
			"rc.rco_start_date",
			"rc.rco_end_date",
			"rc.rco_total_payments",
			"rc.rco_monthly_amount",
			"rc.rco_deposit_amount",
			"rc.rco_currency",
			"rc.rco_due_day",
			"rc.rco_daily_interest_percentage",
			"rc.rco_adjustment_frequency_months",
			"COALESCE(rc.rco_notes, '')",
			"rc.rco_created_at",
			"rc.rco_updated_at",
			"ict.ict_id",
			"ict.ict_code",
			"ict.ict_name",
			"rat.rat_id",
			"rat.rat_code",
			"rat.rat_name",
			"cst.cst_id",
			"cst.cst_code",
			"cst.cst_name",
			"p.pro_id",
			"p.pro_title",
			"t.ten_id",
			"t.ten_name",
		).
		From("rental_contracts rc").
		LeftJoin("interest_calculation_types ict ON ict.ict_id = rc.ict_id").
		LeftJoin("rent_adjustment_types rat ON rat.rat_id = rc.rat_id").
		LeftJoin("contract_statuses cst ON cst.cst_id = rc.cst_id").
		LeftJoin("properties p ON p.pro_id = rc.pro_id").
		LeftJoin("tenants t ON t.ten_id = rc.ten_id")
}

type rentalContractScanner interface {
	Scan(dest ...any) error
}

func scanRentalContract(scanner rentalContractScanner) (domain.RentalContract, error) {
	var rentalContract domain.RentalContract

	var interestCalculationTypeID sql.NullInt64
	var interestCalculationTypeCode sql.NullString
	var interestCalculationTypeName sql.NullString
	var adjustmentTypeID sql.NullInt64
	var adjustmentTypeCode sql.NullString
	var adjustmentTypeName sql.NullString
	var statusID sql.NullInt64
	var statusCode sql.NullString
	var statusName sql.NullString
	var propertyID sql.NullInt64
	var propertyTitle sql.NullString
	var tenantID sql.NullInt64
	var tenantTitle sql.NullString

	err := scanner.Scan(
		&rentalContract.ID,
		&rentalContract.PropertyID,
		&rentalContract.TenantID,
		&rentalContract.StatusID,
		&rentalContract.InterestCalculationTypeID,
		&rentalContract.AdjustmentTypeID,
		&rentalContract.StartDate,
		&rentalContract.EndDate,
		&rentalContract.TotalPayments,
		&rentalContract.MonthlyAmount,
		&rentalContract.DepositAmount,
		&rentalContract.Currency,
		&rentalContract.DueDay,
		&rentalContract.DailyInterestPercentage,
		&rentalContract.AdjustmentFrequencyMonths,
		&rentalContract.Notes,
		&rentalContract.CreatedAt,
		&rentalContract.UpdatedAt,
		&interestCalculationTypeID,
		&interestCalculationTypeCode,
		&interestCalculationTypeName,
		&adjustmentTypeID,
		&adjustmentTypeCode,
		&adjustmentTypeName,
		&statusID,
		&statusCode,
		&statusName,
		&propertyID,
		&propertyTitle,
		&tenantID,
		&tenantTitle,
	)
	if err != nil {
		return rentalContract, err
	}

	if interestCalculationTypeID.Valid {
		rentalContract.InterestCalculationType = &domain.InterestCalculationType{
			ID:   interestCalculationTypeID.Int64,
			Code: interestCalculationTypeCode.String,
			Name: interestCalculationTypeName.String,
		}
	}

	if adjustmentTypeID.Valid {
		rentalContract.AdjustmentType = &domain.RentAdjustmentType{
			ID:   adjustmentTypeID.Int64,
			Code: adjustmentTypeCode.String,
			Name: adjustmentTypeName.String,
		}
	}

	if statusID.Valid {
		rentalContract.Status = &domain.ContractStatus{
			ID:   statusID.Int64,
			Code: statusCode.String,
			Name: statusName.String,
		}
	}

	if propertyID.Valid {
		rentalContract.Property = &domain.Property{ID: propertyID.Int64, Title: propertyTitle.String}
	}

	if tenantID.Valid {
		rentalContract.Tenant = &domain.Tenant{ID: tenantID.Int64, Name: tenantTitle.String}
	}

	return rentalContract, nil
}
