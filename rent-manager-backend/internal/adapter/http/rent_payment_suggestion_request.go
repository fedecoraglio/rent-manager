package http

type getRentPaymentSuggestionRequest struct {
	Period      string `form:"period" binding:"required"`
	PaymentDate string `form:"payment_date" binding:"required"`
}
