package policy

import (
	"regexp"
	"rent-manager-backend/internal/core/domain"
)

var documentID = regexp.MustCompile(`^[A-Za-z0-9_~\-]{1,256}$`)

func ValidateDocumentID(id string) error {
	if !documentID.MatchString(id) {
		return domain.NewAppError(domain.ErrCodeInvalidData, "Invalid document ID")
	}
	return nil
}

func ValidateRentalContractID(id int64) error {
	if id <= 0 {
		return domain.NewAppError(domain.ErrCodeInvalidData, "Invalid rental contract ID")
	}
	return nil
}
