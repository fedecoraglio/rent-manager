package http

type paginationRequest struct {
	Page  uint64 `form:"page" binding:"required,min=1"`
	Limit uint64 `form:"limit" binding:"required,min=1,max=100"`
}
