package domain

import "fmt"

type ErrorCode string

type AppError struct {
	Code        ErrorCode
	Description string
	Cause       error
}

func (e *AppError) Error() string {
	if e == nil {
		return ""
	}

	if e.Cause != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Description, e.Cause)
	}

	return fmt.Sprintf("%s: %s", e.Code, e.Description)
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func NewAppError(code ErrorCode, description string) *AppError {
	return &AppError{
		Code:        code,
		Description: description,
	}
}

func WrapAppError(code ErrorCode, description string, cause error) *AppError {
	return &AppError{
		Code:        code,
		Description: description,
		Cause:       cause,
	}
}
