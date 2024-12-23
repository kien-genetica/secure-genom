package entity

import "time"

type UploadStatus string

const (
	UploadStatusPending  UploadStatus = "pending"
	UploadStatusComplete UploadStatus = "complete"
	UploadStatusFailed   UploadStatus = "failed"
)

type Data struct {
	ID              string
	OwnerID         string
	EncryptedData   []byte
	PublicKey       string
	ValidationState string
	ProcessedState  string
	CreatedAt       time.Time
}
