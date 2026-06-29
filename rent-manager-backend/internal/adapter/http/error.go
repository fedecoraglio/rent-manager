package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"rent-manager-backend/internal/core/domain"
)

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func validationError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, errorResponse{
		Code:    string(domain.ErrCodeInvalidData),
		Message: err.Error(),
	})
}

func handleError(ctx *gin.Context, err error) {
	var appErr *domain.AppError

	if errors.As(err, &appErr) {
		ctx.JSON(statusFromErrorCode(appErr.Code), errorResponse{
			Code:    string(appErr.Code),
			Message: appErr.Description,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, errorResponse{
		Code:    "CORE_9999",
		Message: "internal server error",
	})
}

func statusFromErrorCode(code domain.ErrorCode) int {
	switch code {
	case domain.ErrCodeDataNotFound, domain.ErrCodeUserNotFound, domain.ErrCodePropertyNotFound:
		return http.StatusNotFound
	case domain.ErrCodeConflictingData, domain.ErrCodeUserAlreadyExists, domain.ErrCodePropertyAlreadyExists:
		return http.StatusConflict
	case domain.ErrCodeInvalidData,
		domain.ErrCodeUserNil,
		domain.ErrCodeUserNameEmpty,
		domain.ErrCodeUserEmailEmpty,
		domain.ErrCodeUserEmailInvalid,
		domain.ErrCodeUserPasswordEmpty:
		return http.StatusBadRequest
	case domain.ErrCodeOwnerNotFound:
		return http.StatusNotFound
	case domain.ErrCodeOwnerAlreadyExists:
		return http.StatusConflict
	case domain.ErrCodeOwnerNil,
		domain.ErrCodeOwnerNameEmpty,
		domain.ErrCodeOwnerDocumentNumberEmpty,
		domain.ErrCodeOwnerInvalidEmail,
		domain.ErrCodePropertyOwnerIDEmpty,
		domain.ErrCodePropertyTypeIDEmpty,
		domain.ErrCodePropertyStatusIDEmpty,
		domain.ErrCodePropertyCountryIDEmpty,
		domain.ErrCodePropertyStateIDEmpty,
		domain.ErrCodePropertyCodeEmpty,
		domain.ErrCodePropertyTitleEmpty,
		domain.ErrCodePropertyStreetEmpty,
		domain.ErrCodePropertyCityEmpty:
		return http.StatusBadRequest
	case domain.ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case domain.ErrCodeForbidden:
		return http.StatusForbidden
	case domain.ErrCodeRentalContractNotFound:
		return http.StatusNotFound
	case domain.ErrCodeRentalContractAlreadyExists:
		return http.StatusConflict
	case domain.ErrCodeRentalContractNil,
		domain.ErrCodeRentalContractPropertyIDEmpty,
		domain.ErrCodeRentalContractTenantIDEmpty,
		domain.ErrCodeRentalContractStatusIDEmpty,
		domain.ErrCodeRentalContractStartDateEmpty,
		domain.ErrCodeRentalContractEndDateEmpty,
		domain.ErrCodeRentalContractInvalidDateRange,
		domain.ErrCodeRentalContractMonthlyAmountInvalid,
		domain.ErrCodeRentalContractCurrencyEmpty,
		domain.ErrCodeRentalContractInterestCalculationTypeIDEmpty,
		domain.ErrCodeRentalContractAdjustmentTypeIDEmpty,
		domain.ErrCodeRentalContractDepositAmountInvalid,
		domain.ErrCodeRentalContractDueDayInvalid,
		domain.ErrCodeRentalContractDailyInterestPercentageInvalid,
		domain.ErrCodeRentalContractAdjustmentFrequencyInvalid:
		return http.StatusBadRequest
	case domain.ErrCodeRentPaymentNotFound:
		return http.StatusNotFound
	case domain.ErrCodeRentPaymentAlreadyExists:
		return http.StatusConflict
	case domain.ErrCodeRentPaymentNil,
		domain.ErrCodeRentPaymentContractIDEmpty,
		domain.ErrCodeRentPaymentPeriodEmpty,
		domain.ErrCodeRentPaymentDueDateEmpty,
		domain.ErrCodeRentPaymentBaseAmountInvalid,
		domain.ErrCodeRentPaymentTotalAmountInvalid,
		domain.ErrCodeRentPaymentPaidAmountInvalid,
		domain.ErrCodeRentPaymentPeriodAlreadyPaid,
		domain.ErrCodeRentPaymentPeriodOutsideContract:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
