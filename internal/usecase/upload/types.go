package upload

import "context"

type UploadUseCase interface {
	UploadData(ctx context.Context, data []byte) error
}
