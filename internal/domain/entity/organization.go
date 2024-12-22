package entity

import "time"

type Organization struct {
	ID           string
	Name         string
	PublicKey    string
	ReEncryptKey string
	IsAuthorized bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
