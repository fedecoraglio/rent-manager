package domain

import "errors"

var (
	ErrInflationIndexNotFound          = errors.New("inflation index not found")
	ErrInflationIndexPeriodRequired    = errors.New("inflation index period is required")
	ErrInflationIndexPercentageInvalid = errors.New("inflation index percentage is invalid")
	ErrInflationIndexAlreadyExists     = errors.New("inflation index already exists")
)
