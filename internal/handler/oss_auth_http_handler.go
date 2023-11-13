package handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/model"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/internal/usecase"
	pkgError "gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/pkg/error"
	"gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/pkg/helper"
	pkgValidator "gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/pkg/validator"
)

type ossAuthHttpHandler struct {
	ossAuthUsecase usecase.OssAuthUsecase
	validator      *pkgValidator.XValidator
}

func InitOssAuthHttpHandler(app *fiber.App, ossAuthUsecase usecase.OssAuthUsecase, validator *pkgValidator.XValidator) {
	handler := &ossAuthHttpHandler{
		ossAuthUsecase: ossAuthUsecase,
		validator:      validator,
	}

	route := app.Group("/oss-auth")

	route.Get("/", handler.Get)
}

func (p *ossAuthHttpHandler) Get(c *fiber.Ctx) error {
	request := &model.GetRequest{}

	if err := c.QueryParser(request); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}

	if err := p.validator.Validate(request); err != nil {
		return helper.RespError(c, err)
	}

	response, err := p.ossAuthUsecase.Get(c.Context(), *request)
	if err != nil {
		return helper.RespError(c, err)
	}

	return helper.RespSuccess(c, response, "get success")
}
