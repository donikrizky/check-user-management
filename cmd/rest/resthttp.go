package rest

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/config"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/handler"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/model"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/repository"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/usecase"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/pkg/db"
	pkgValidator "gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/pkg/validator"
)

func ServeREST() error {
	config, err := config.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg("load config fail")
	}

	fmt.Println(config)

	db, err := db.NewDb(
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.OssAuth{})

	ossAuthRepository := repository.NewOssAuthRepository(db)
	ossAuthUsecase := usecase.NewOssAuthRepository(ossAuthRepository)

	app := fiber.New()

	handler.InitOssAuthHttpHandler(app, ossAuthUsecase, &pkgValidator.XValidator{
		Validator: &validator.Validate{},
	})

	app.Listen(fmt.Sprintf(":%s", config.AppPort))

	return nil
}
