package repository

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"log/slog"
	"rent-manager-backend/internal/adapter/store/mysql"
	"rent-manager-backend/internal/core/domain"
	"time"
)

type RentPaymentRepository struct {
	db *mysql.DB
}

func (rpr *RentPaymentRepository) CreateRentPayment(
	ctx context.Context,
	rentPayment *domain.RentPayment,
) (*domain.RentPayment, error) {
	query := rpr.db.QueryBuilder.
		Insert("rent_payments").
		Columns(
			"rco_id",
			"rpa_period",
			"rpa_due_date",
			"rpa_payment_date",
			"rpa_base_amount",
			"rpa_suggested_adjustment_percentage",
			"rpa_applied_adjustment_percentage",
			"rpa_suggested_interest_amount",
			"rpa_applied_interest_amount",
			"rpa_total_amount",
			"rpa_paid_amount",
			"rpa_is_paid",
			"rpa_notes",
		).
		Values(
			rentPayment.RentalContractID,
			rentPayment.Period,
			rentPayment.DueDate,
			rentPayment.PaymentDate,
			rentPayment.BaseAmount,
			rentPayment.SuggestedAdjustmentPercentage,
			rentPayment.AppliedAdjustmentPercentage,
			rentPayment.SuggestedInterestAmount,
			rentPayment.AppliedInterestAmount,
			rentPayment.TotalAmount,
			rentPayment.PaidAmount,
			rentPayment.IsPaid,
			rentPayment.Notes,
		)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[CreateRentPayment] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := rpr.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[CreateRentPayment] failed to execute sql query", "err", err)

		if errCode := rpr.db.ErrorCode(err); errCode == "1062" {
			return nil, domain.ErrRentPaymentAlreadyExists
		}

		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		slog.Error("[CreateRentPayment] failed to get last insert id", "err", err)
		return nil, err
	}

	return rpr.GetRentPaymentByID(ctx, lastID)
}

func (rpr *RentPaymentRepository) UpdateRentPayment(
	ctx context.Context,
	rentPayment *domain.RentPayment,
) (*domain.RentPayment, error) {
	queryBuilder := rpr.db.QueryBuilder.Update("rent_payments")

	if rentPayment.RentalContractID != 0 {
		queryBuilder = queryBuilder.Set("rco_id", rentPayment.RentalContractID)
	}

	if !rentPayment.Period.IsZero() {
		queryBuilder = queryBuilder.Set("rpa_period", rentPayment.Period)
	}

	if !rentPayment.DueDate.IsZero() {
		queryBuilder = queryBuilder.Set("rpa_due_date", rentPayment.DueDate)
	}

	if rentPayment.PaymentDate != nil {
		queryBuilder = queryBuilder.Set("rpa_payment_date", rentPayment.PaymentDate)
	}

	if rentPayment.BaseAmount != 0 {
		queryBuilder = queryBuilder.Set("rpa_base_amount", rentPayment.BaseAmount)
	}

	queryBuilder = queryBuilder.Set(
		"rpa_suggested_adjustment_percentage",
		rentPayment.SuggestedAdjustmentPercentage,
	)

	queryBuilder = queryBuilder.Set(
		"rpa_applied_adjustment_percentage",
		rentPayment.AppliedAdjustmentPercentage,
	)

	queryBuilder = queryBuilder.Set(
		"rpa_suggested_interest_amount",
		rentPayment.SuggestedInterestAmount,
	)

	queryBuilder = queryBuilder.Set(
		"rpa_applied_interest_amount",
		rentPayment.AppliedInterestAmount,
	)

	if rentPayment.TotalAmount != 0 {
		queryBuilder = queryBuilder.Set("rpa_total_amount", rentPayment.TotalAmount)
	}

	queryBuilder = queryBuilder.Set("rpa_paid_amount", rentPayment.PaidAmount)
	queryBuilder = queryBuilder.Set("rpa_is_paid", rentPayment.IsPaid)
	queryBuilder = queryBuilder.Set("rpa_notes", rentPayment.Notes)

	queryBuilder = queryBuilder.Where(sq.Eq{"rpa_id": rentPayment.ID})

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		slog.Error("[UpdateRentPayment] failed to generate sql query", "err", err)
		return nil, err
	}

	result, err := rpr.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[UpdateRentPayment] failed to execute sql query", "err", err)

		if errCode := rpr.db.ErrorCode(err); errCode == "1062" {
			return nil, domain.ErrRentPaymentAlreadyExists
		}

		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("[UpdateRentPayment] failed to get rows affected", "err", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domain.ErrDataNotFound
	}

	return rpr.GetRentPaymentByID(ctx, rentPayment.ID)
}

func (rpr *RentPaymentRepository) GetRentPaymentByID(
	ctx context.Context,
	id int64,
) (*domain.RentPayment, error) {
	var rentPayment domain.RentPayment
	var paymentDate sql.NullTime

	query := rpr.baseRentPaymentQuery().
		Where(sq.Eq{"rpa.rpa_id": id}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetRentPaymentByID] failed to generate sql query", "err", err)
		return nil, err
	}

	err = rpr.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&rentPayment.ID,
		&rentPayment.RentalContractID,
		&rentPayment.Period,
		&rentPayment.DueDate,
		&paymentDate,
		&rentPayment.BaseAmount,
		&rentPayment.SuggestedAdjustmentPercentage,
		&rentPayment.AppliedAdjustmentPercentage,
		&rentPayment.SuggestedInterestAmount,
		&rentPayment.AppliedInterestAmount,
		&rentPayment.TotalAmount,
		&rentPayment.PaidAmount,
		&rentPayment.IsPaid,
		&rentPayment.Notes,
		&rentPayment.CreatedAt,
		&rentPayment.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrDataNotFound
		}

		slog.Error("[GetRentPaymentByID] failed to scan row", "err", err)
		return nil, err
	}

	if paymentDate.Valid {
		rentPayment.PaymentDate = &paymentDate.Time
	}

	return &rentPayment, nil
}

func (rpr *RentPaymentRepository) GetRentPaymentByContractIDAndPeriod(
	ctx context.Context,
	rentalContractID int64,
	period time.Time,
) (*domain.RentPayment, error) {
	var rentPayment domain.RentPayment
	var paymentDate sql.NullTime

	query := rpr.baseRentPaymentQuery().
		Where(sq.Eq{"rpa.rco_id": rentalContractID}).
		Where(sq.Eq{"rpa.rpa_period": period}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[GetRentPaymentByContractIDAndPeriod] failed to generate sql query", "err", err)
		return nil, err
	}

	err = rpr.db.QueryRowContext(ctx, sqlQuery, args...).Scan(
		&rentPayment.ID,
		&rentPayment.RentalContractID,
		&rentPayment.Period,
		&rentPayment.DueDate,
		&paymentDate,
		&rentPayment.BaseAmount,
		&rentPayment.SuggestedAdjustmentPercentage,
		&rentPayment.AppliedAdjustmentPercentage,
		&rentPayment.SuggestedInterestAmount,
		&rentPayment.AppliedInterestAmount,
		&rentPayment.TotalAmount,
		&rentPayment.PaidAmount,
		&rentPayment.IsPaid,
		&rentPayment.Notes,
		&rentPayment.CreatedAt,
		&rentPayment.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		slog.Error("[GetRentPaymentByContractIDAndPeriod] failed to scan row", "err", err)
		return nil, err
	}

	if paymentDate.Valid {
		rentPayment.PaymentDate = &paymentDate.Time
	}

	return &rentPayment, nil
}
func (rpr *RentPaymentRepository) ListRentPayments(
	ctx context.Context,
	rentalContractID int64,
	page uint64,
	limit uint64,
) ([]domain.RentPayment, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := rpr.baseRentPaymentQuery().
		OrderBy("rpa.rpa_period").
		Limit(limit).
		Offset(offset)

	if rentalContractID > 0 {
		query = query.Where(sq.Eq{"rpa.rco_id": rentalContractID})
	}

	return rpr.findRentPayments(ctx, query)
}

func NewRentPaymentRepository(db *mysql.DB) *RentPaymentRepository {
	return &RentPaymentRepository{
		db: db,
	}
}

func (rpr *RentPaymentRepository) ListRentPaymentsByContractID(
	ctx context.Context,
	rentalContractID int64,
) ([]domain.RentPayment, error) {
	query := rpr.baseRentPaymentQuery().
		Where(sq.Eq{"rpa.rco_id": rentalContractID}).
		OrderBy("rpa.rpa_period")

	return rpr.findRentPayments(ctx, query)
}

func (rpr *RentPaymentRepository) baseRentPaymentQuery() sq.SelectBuilder {
	return rpr.db.QueryBuilder.
		Select(
			"rpa.rpa_id",
			"rpa.rco_id",
			"rpa.rpa_period",
			"rpa.rpa_due_date",
			"rpa.rpa_payment_date",
			"rpa.rpa_base_amount",
			"COALESCE(rpa.rpa_suggested_adjustment_percentage, 0)",
			"COALESCE(rpa.rpa_applied_adjustment_percentage, 0)",
			"COALESCE(rpa.rpa_suggested_interest_amount, 0)",
			"COALESCE(rpa.rpa_applied_interest_amount, 0)",
			"rpa.rpa_total_amount",
			"COALESCE(rpa.rpa_paid_amount, 0)",
			"rpa.rpa_is_paid",
			"COALESCE(rpa.rpa_notes, '')",
			"rpa.rpa_created_at",
			"rpa.rpa_updated_at",
		).
		From("rent_payments rpa")
}

func (rpr *RentPaymentRepository) findRentPayments(
	ctx context.Context,
	query sq.SelectBuilder,
) ([]domain.RentPayment, error) {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		slog.Error("[findRentPayments] failed to generate sql query", "err", err)
		return nil, err
	}

	rows, err := rpr.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		slog.Error("[findRentPayments] failed to execute sql query", "err", err)
		return nil, err
	}
	defer rows.Close()

	var rentPayments []domain.RentPayment

	for rows.Next() {
		var rentPayment domain.RentPayment
		var paymentDate sql.NullTime

		err := rows.Scan(
			&rentPayment.ID,
			&rentPayment.RentalContractID,
			&rentPayment.Period,
			&rentPayment.DueDate,
			&paymentDate,
			&rentPayment.BaseAmount,
			&rentPayment.SuggestedAdjustmentPercentage,
			&rentPayment.AppliedAdjustmentPercentage,
			&rentPayment.SuggestedInterestAmount,
			&rentPayment.AppliedInterestAmount,
			&rentPayment.TotalAmount,
			&rentPayment.PaidAmount,
			&rentPayment.IsPaid,
			&rentPayment.Notes,
			&rentPayment.CreatedAt,
			&rentPayment.UpdatedAt,
		)
		if err != nil {
			slog.Error("[findRentPayments] failed to scan row", "err", err)
			return nil, err
		}

		if paymentDate.Valid {
			rentPayment.PaymentDate = &paymentDate.Time
		}

		rentPayments = append(rentPayments, rentPayment)
	}

	if err := rows.Err(); err != nil {
		slog.Error("[findRentPayments] rows iteration failed", "err", err)
		return nil, err
	}

	return rentPayments, nil
}
