package http

type InflationIndexCreateRequest struct {
	Period     string  `json:"period" binding:"required"` // yyyy-mm
	Percentage float64 `json:"percentage" binding:"required"`
	Source     string  `json:"source"`
	Notes      string  `json:"notes"`
}
type InflationIndexUpdateRequest struct {
	Period     string  `json:"period" binding:"required"` // yyyy-mm
	Percentage float64 `json:"percentage" binding:"required"`
	Source     string  `json:"source"`
	Notes      string  `json:"notes"`
}
