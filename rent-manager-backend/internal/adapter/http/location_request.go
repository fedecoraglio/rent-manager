package http

type listStatesByCountryRequest struct {
	CountryID int64 `uri:"countryId" binding:"required,min=1"`
}
