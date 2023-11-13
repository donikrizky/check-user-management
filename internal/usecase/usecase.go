package usecase

import (
	"context"

	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/model"
)

type OssAuthUsecase interface {
	Get(ctx context.Context, payload model.GetRequest) (*model.GetResponse, error)
}
