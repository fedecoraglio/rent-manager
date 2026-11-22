package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"rent-manager-backend/internal/core/domain"
	documentUseCase "rent-manager-backend/internal/core/usecase/document"
	"strconv"
	"time"
)

type DocumentHandler struct {
	uploadDocumentUseCase      *documentUseCase.UploadDocumentUseCase
	listDocumentsUseCase       *documentUseCase.ListDocumentsUseCase
	downloadDocumentUseCase    *documentUseCase.DownloadDocumentUseCase
	deleteDocumentUseCase      *documentUseCase.DeleteDocumentUseCase
	askContractQuestionUseCase *documentUseCase.AskContractQuestionUseCase
}

func NewDocumentHandler(
	uploadDocumentUseCase *documentUseCase.UploadDocumentUseCase,
	listDocumentsUseCase *documentUseCase.ListDocumentsUseCase,
	downloadDocumentUseCase *documentUseCase.DownloadDocumentUseCase,
	deleteDocumentUseCase *documentUseCase.DeleteDocumentUseCase,
	askContractQuestionUseCase *documentUseCase.AskContractQuestionUseCase,
) *DocumentHandler {
	return &DocumentHandler{
		uploadDocumentUseCase:      uploadDocumentUseCase,
		listDocumentsUseCase:       listDocumentsUseCase,
		downloadDocumentUseCase:    downloadDocumentUseCase,
		deleteDocumentUseCase:      deleteDocumentUseCase,
		askContractQuestionUseCase: askContractQuestionUseCase,
	}
}
func (h *DocumentHandler) Register(routes *gin.RouterGroup) {
	routes.POST("/:rentalContractId/documents", h.Upload)
	routes.GET("/:rentalContractId/documents", h.List)
	routes.GET("/:rentalContractId/documents/:documentId", h.Download)
	routes.DELETE("/:rentalContractId/documents/:documentId", h.Delete)
	routes.POST("/:rentalContractId/chat", h.Ask)
}

type documentResponse struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Filename    string    `json:"filename"`
	ContentType string    `json:"contentType"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"createdAt"`
}

func newDocumentResponse(d domain.Document) documentResponse {
	return documentResponse{d.ID, d.Type, d.Filename, d.ContentType, d.Size, d.CreatedAt}
}
func documentContractID(ctx *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(ctx.Param("rentalContractId"), 10, 64)
	if err != nil || id <= 0 {
		validationError(ctx, fmt.Errorf("invalid rental contract ID"))
		return 0, false
	}
	return id, true
}
func (h *DocumentHandler) Upload(ctx *gin.Context) {
	id, ok := documentContractID(ctx)
	if !ok {
		return
	}
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 21*1024*1024)
	if err := ctx.Request.ParseMultipartForm(1024 * 1024); err != nil {
		if ctx.Request.MultipartForm != nil {
			defer ctx.Request.MultipartForm.RemoveAll()
		}
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			handleError(ctx, domain.NewAppError(domain.ErrCodeDocumentTooLarge, "Multipart body exceeds 21 MiB"))
		} else {
			validationError(ctx, fmt.Errorf("invalid multipart upload"))
		}
		return
	}
	defer ctx.Request.MultipartForm.RemoveAll()
	files := ctx.Request.MultipartForm.File["file"]
	types := ctx.Request.MultipartForm.Value["type"]
	if len(files) != 1 || len(types) != 1 {
		validationError(ctx, fmt.Errorf("provide exactly one file and document type"))
		return
	}
	file := files[0]
	contentType, _, err := mime.ParseMediaType(file.Header.Get("Content-Type"))
	if err != nil || contentType != "application/pdf" {
		handleError(ctx, domain.NewAppError(domain.ErrCodeDocumentUnsupported, "Only application/pdf is supported"))
		return
	}
	input, err := file.Open()
	if err != nil {
		handleError(ctx, err)
		return
	}
	defer input.Close()
	result, err := h.uploadDocumentUseCase.Execute(ctx.Request.Context(), id, domain.DocumentUpload{Filename: file.Filename, ContentType: contentType, Type: types[0], Size: file.Size, Content: input})
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.Header("Location", fmt.Sprintf("/v1/rental-contracts/%d/documents/%s", id, result.ID))
	ctx.JSON(http.StatusCreated, successResponse{Data: newDocumentResponse(*result)})
}
func (h *DocumentHandler) List(ctx *gin.Context) {
	id, ok := documentContractID(ctx)
	if !ok {
		return
	}
	result, err := h.listDocumentsUseCase.Execute(ctx.Request.Context(), id)
	if err != nil {
		handleError(ctx, err)
		return
	}
	documents := make([]documentResponse, 0, len(result))
	for _, item := range result {
		documents = append(documents, newDocumentResponse(item))
	}
	handleSuccess(ctx, documents)
}

func (h *DocumentHandler) Download(ctx *gin.Context) {
	id, ok := documentContractID(ctx)
	if !ok {
		return
	}
	attachment := false
	if value, exists := ctx.GetQuery("download"); exists {
		var err error
		attachment, err = strconv.ParseBool(value)
		if err != nil {
			validationError(ctx, fmt.Errorf("invalid download flag"))
			return
		}
	}
	result, err := h.downloadDocumentUseCase.Execute(ctx.Request.Context(), id, ctx.Param("documentId"), attachment)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.Header("Cache-Control", "no-store")
	ctx.Header("X-Content-Type-Options", "nosniff")
	ctx.Header("Content-Disposition", result.ContentDisposition)
	ctx.Data(http.StatusOK, "application/pdf", result.Content)
}
func (h *DocumentHandler) Delete(ctx *gin.Context) {
	id, ok := documentContractID(ctx)
	if !ok {
		return
	}
	if err := h.deleteDocumentUseCase.Execute(ctx.Request.Context(), id, ctx.Param("documentId")); err != nil {
		handleError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (h *DocumentHandler) Ask(ctx *gin.Context) {
	id, ok := documentContractID(ctx)
	if !ok {
		return
	}
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 64*1024)
	var request struct {
		Question string `json:"question"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, fmt.Errorf("invalid chat request"))
		return
	}
	answer, err := h.askContractQuestionUseCase.Execute(ctx.Request.Context(), id, request.Question)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, gin.H{"answer": answer})
}
