package usecase

import (
	"context"
	"fmt"

	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/model"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/repository"
)

type ossAuthUsecase struct {
	ossAuthRepository repository.OssAuthRepository
}

func NewOssAuthRepository(ossAuthRepository repository.OssAuthRepository) OssAuthUsecase {
	return &ossAuthUsecase{
		ossAuthRepository: ossAuthRepository,
	}
}

func (u *ossAuthUsecase) Get(ctx context.Context, payload model.GetRequest) (*model.GetResponse, error) {
	ossAuths, err := u.ossAuthRepository.GetOssAuths(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(ossAuths)

	return &model.GetResponse{
		Test: "test",
	}, nil
}
