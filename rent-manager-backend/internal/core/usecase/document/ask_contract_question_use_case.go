package document

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	documentPolicy "rent-manager-backend/internal/core/policy/document"
	"rent-manager-backend/internal/core/port"
)

type AskContractQuestionUseCase struct {
	contracts      port.RentalContractRepository
	documents      port.DocumentGateway
	questionPolicy *documentPolicy.ContractQuestionPolicy
}

func NewAskContractQuestionUseCase(contracts port.RentalContractRepository, documents port.DocumentGateway, questionPolicy *documentPolicy.ContractQuestionPolicy) *AskContractQuestionUseCase {
	return &AskContractQuestionUseCase{contracts: contracts, documents: documents, questionPolicy: questionPolicy}
}

func (uc *AskContractQuestionUseCase) Execute(ctx context.Context, id int64, question string) (string, error) {
	request := domain.ContractQuestion{Question: question}
	if err := uc.questionPolicy.Execute(ctx, &request); err != nil {
		return "", err
	}
	scope, err := resolveDocumentScope(ctx, uc.contracts, id)
	if err != nil {
		return "", err
	}
	return uc.documents.Ask(ctx, scope, request.Question)
}
