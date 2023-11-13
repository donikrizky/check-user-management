package repository

import (
	"context"

	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/model"
)

type OssAuthRepository interface {
	GetOssAuths(ctx context.Context) ([]model.OssAuth, error)
}
