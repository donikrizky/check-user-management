package repository

import (
	"context"

	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/model"
	"gorm.io/gorm"
)

type ossAuthRepository struct {
	db *gorm.DB
}

func NewOssAuthRepository(db *gorm.DB) OssAuthRepository {
	return &ossAuthRepository{
		db: db,
	}
}

func (r *ossAuthRepository) GetOssAuths(ctx context.Context) ([]model.OssAuth, error) {
	var ossAuths []model.OssAuth
	r.db.Find(&ossAuths)

	return ossAuths, nil
}
