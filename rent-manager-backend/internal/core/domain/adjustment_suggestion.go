package domain

const (
	AdjustmentStatusCalculated     = "calculated"
	AdjustmentStatusNotApplicable  = "not_applicable"
	AdjustmentStatusMissingIndexes = "missing_indexes"
)

type AdjustmentSuggestion struct {
	Percentage float64
	Status     string
	Message    string
}
