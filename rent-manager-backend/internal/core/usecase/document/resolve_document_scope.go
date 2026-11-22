package document

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	documentPolicy "rent-manager-backend/internal/core/policy/document"
	"rent-manager-backend/internal/core/port"
)

// Scope resolution follows existing administrative JWT access. The data model
// does not associate authenticated users with individual property owners.
func resolveDocumentScope(ctx context.Context, contracts port.RentalContractRepository, id int64) (domain.DocumentScope, error) {
	if err := documentPolicy.ValidateRentalContractID(id); err != nil {
		return domain.DocumentScope{}, err
	}
	contract, err := contracts.GetRentalContractByID(ctx, id)
	if err != nil {
		return domain.DocumentScope{}, err
	}
	if contract == nil || contract.PropertyID <= 0 {
		return domain.DocumentScope{}, domain.ErrRentalContractNotFound
	}
	return domain.DocumentScope{PropertyID: contract.PropertyID, ContractID: contract.ID}, nil
}
