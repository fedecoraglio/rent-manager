package http

import (
	"rent-manager-backend/internal/core/domain"
	"time"
)

type InflationIndexResponse struct {
	ID         int64   `json:"id"`
	Period     string  `json:"period"`
	Percentage float64 `json:"percentage"`
	Source     string  `json:"source"`
	Notes      string  `json:"notes"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

func newInflationIndexResponse(inflationIndex *domain.InflationIndex) InflationIndexResponse {
	return InflationIndexResponse{
		ID:         inflationIndex.ID,
		Period:     inflationIndex.Period.Format("2006-01"),
		Percentage: inflationIndex.Percentage,
		Source:     inflationIndex.Source,
		Notes:      inflationIndex.Notes,
		CreatedAt:  inflationIndex.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  inflationIndex.UpdatedAt.Format(time.RFC3339),
	}
}

func newInflationIndexesResponse(inflationIndexes []domain.InflationIndex) []InflationIndexResponse {
	response := make([]InflationIndexResponse, 0, len(inflationIndexes))

	for _, inflationIndex := range inflationIndexes {
		inflationIndexCopy := inflationIndex
		response = append(response, newInflationIndexResponse(&inflationIndexCopy))
	}

	return response
}
