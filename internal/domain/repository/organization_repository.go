package repository

import (
	"context"

	"github.com/kien6034/secure-genom/internal/domain/entity"
)

type OrganizationRepository interface {
	Create(ctx context.Context, org *entity.Organization) error
	FindByID(ctx context.Context, id string) (*entity.Organization, error)
	Update(ctx context.Context, org *entity.Organization) error
}
