package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	locationUseCase "rent-manager-backend/internal/core/usecase/location"
)

type LocationHandler struct {
	listCountriesUseCase       *locationUseCase.ListCountriesUseCase
	listStatesByCountryUseCase *locationUseCase.ListStatesByCountryUseCase
}

func NewLocationHandler(
	listCountriesUseCase *locationUseCase.ListCountriesUseCase,
	listStatesByCountryUseCase *locationUseCase.ListStatesByCountryUseCase,
) *LocationHandler {
	return &LocationHandler{
		listCountriesUseCase:       listCountriesUseCase,
		listStatesByCountryUseCase: listStatesByCountryUseCase,
	}
}

func (h *LocationHandler) ListCountries(ctx *gin.Context) {
	countries, err := h.listCountriesUseCase.ListCountries(ctx)
	if err != nil {
		slog.Error("[ListCountries] failed to list countries", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newCountriesResponse(countries))
}

func (h *LocationHandler) ListStatesByCountry(ctx *gin.Context) {
	var req listStatesByCountryRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	states, err := h.listStatesByCountryUseCase.ListStatesByCountry(
		ctx,
		req.CountryID,
	)
	if err != nil {
		slog.Error("[ListStatesByCountry] failed to list states by country", "err", err)
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newStatesResponse(states))
}
