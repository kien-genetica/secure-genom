package entity

import "time"

type Data struct {
	ID              string
	OwnerID         string
	EncryptedData   []byte
	PublicKey       string
	ValidationState string
	ProcessedState  string
	CreatedAt       time.Time
}
