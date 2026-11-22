package policy

import (
	"context"
	"rent-manager-backend/internal/core/domain"
	"strings"
)

type ContractQuestionPolicy struct{}

func NewContractQuestionPolicy() *ContractQuestionPolicy { return &ContractQuestionPolicy{} }

func (policy *ContractQuestionPolicy) Execute(_ context.Context, question *domain.ContractQuestion) error {
	if question == nil {
		return domain.NewAppError(domain.ErrCodeInvalidData, "Question must contain 1 to 8000 bytes")
	}
	question.Question = strings.TrimSpace(question.Question)
	if len(question.Question) == 0 || len(question.Question) > 8000 {
		return domain.NewAppError(domain.ErrCodeInvalidData, "Question must contain 1 to 8000 bytes")
	}
	return nil
}
