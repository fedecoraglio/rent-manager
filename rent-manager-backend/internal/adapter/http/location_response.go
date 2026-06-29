package http

import (
	"time"

	"rent-manager-backend/internal/core/domain"
)

type countryResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type stateResponse struct {
	ID        int64     `json:"id"`
	CountryID int64     `json:"country_id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newCountryResponse(country domain.Country) countryResponse {
	return countryResponse{
		ID:        country.ID,
		Code:      country.Code,
		Name:      country.Name,
		CreatedAt: country.CreatedAt,
		UpdatedAt: country.UpdatedAt,
	}
}

func newCountriesResponse(countries []domain.Country) []countryResponse {
	response := make([]countryResponse, 0, len(countries))

	for _, country := range countries {
		response = append(response, newCountryResponse(country))
	}

	return response
}

func newStateResponse(state domain.State) stateResponse {
	return stateResponse{
		ID:        state.ID,
		CountryID: state.CountryID,
		Code:      state.Code,
		Name:      state.Name,
		CreatedAt: state.CreatedAt,
		UpdatedAt: state.UpdatedAt,
	}
}

func newStatesResponse(states []domain.State) []stateResponse {
	response := make([]stateResponse, 0, len(states))

	for _, state := range states {
		response = append(response, newStateResponse(state))
	}

	return response
}
