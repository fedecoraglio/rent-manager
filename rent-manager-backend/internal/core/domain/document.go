package domain

import (
	"io"
	"time"
)

const MaxDocumentBytes int64 = 20 * 1024 * 1024

type DocumentScope struct{ PropertyID, ContractID int64 }

type ContractQuestion struct{ Question string }
type Document struct {
	ID, Type, Filename, ContentType string
	Size                            int64
	CreatedAt                       time.Time
}
type DocumentUpload struct {
	Filename, ContentType, Type string
	Size                        int64
	Content                     io.Reader
}
type DocumentDownload struct {
	Content            []byte
	ContentDisposition string
}

const (
	ErrCodeAIUnavailable         ErrorCode = "AI_0001"
	ErrCodeAITimeout             ErrorCode = "AI_0002"
	ErrCodeDocumentTooLarge      ErrorCode = "AI_0003"
	ErrCodeDocumentUnsupported   ErrorCode = "AI_0004"
	ErrCodeDocumentUnprocessable ErrorCode = "AI_0005"
)
