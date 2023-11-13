package helper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	pkgError "gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/pkg/error"
)

// Result common output
type Result struct {
	Data     interface{}
	MetaData interface{}
	Error    error
	Count    int64
}

type response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MetaData struct {
	Page      int64 `json:"page"`
	Count     int64 `json:"count"`
	TotalPage int64 `json:"totalPage"`
	TotalData int64 `json:"totalData"`
}

type paginationResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    MetaData    `json:"meta"`
}

func getErrorStatusCode(err error) int {
	errString, ok := err.(*pkgError.ErrorString)
	if ok {
		return errString.Code()
	}

	// default http status code
	return http.StatusInternalServerError
}

type Meta struct {
	Method        string    `json:"method"`
	Url           string    `json:"url"`
	Code          string    `json:"code"`
	ContentLength int64     `json:"content_length"`
	Date          time.Time `json:"date"`
	Ip            string    `json:"ip"`
}

func RespSuccess(c *fiber.Ctx, data interface{}, message string) error {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.Path(),
		Method:        c.Method(),
		Code:          fmt.Sprintf("%v", http.StatusOK),
		Ip:            c.IP(),
		ContentLength: int64(c.Request().Header.ContentLength()),
	}
	log.Info().Interface("meta", meta).Msg(message)

	return c.Status(http.StatusOK).JSON(response{
		Message: message,
		Data:    data,
		Code:    http.StatusOK,
		Success: true,
	})
}

func RespError(c *fiber.Ctx, err error) error {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.Path(),
		Method:        c.Method(),
		Code:          fmt.Sprintf("%v", getErrorStatusCode(err)),
		Ip:            c.IP(),
		ContentLength: int64(c.Request().Header.ContentLength()),
	}
	log.Info().Interface("meta", meta).Msg(err.Error())

	return c.Status(getErrorStatusCode(err)).JSON(response{
		Message: err.Error(),
		Data:    nil,
		Code:    getErrorStatusCode(err),
		Success: false,
	})
}

func RespPagination(c *fiber.Ctx, data interface{}, metadata MetaData, message string) error {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.Path(),
		Method:        c.Method(),
		Code:          fmt.Sprintf("%v", http.StatusOK),
		Ip:            c.IP(),
		ContentLength: int64(c.Request().Header.ContentLength()),
	}
	log.Info().Interface("meta", meta).Msg(message)

	return c.Status(http.StatusOK).JSON(paginationResponse{
		Message: message,
		Meta:    metadata,
		Data:    data,
		Code:    http.StatusOK,
		Success: true,
	})
}

func RespErrorWithData(c *fiber.Ctx, data interface{}, err error) error {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.Path(),
		Method:        c.Method(),
		Code:          fmt.Sprintf("%v", getErrorStatusCode(err)),
		Ip:            c.IP(),
		ContentLength: int64(c.Request().Header.ContentLength()),
	}
	log.Info().Interface("meta", meta).Msg(err.Error())

	return c.Status(getErrorStatusCode(err)).JSON(response{
		Message: err.Error(),
		Data:    data,
		Code:    getErrorStatusCode(err),
		Success: false,
	})
}

func RespCustomError(c *fiber.Ctx, err error) error {
	meta := Meta{
		Date:          time.Now(),
		Url:           c.Path(),
		Method:        c.Method(),
		Code:          fmt.Sprintf("%v", getErrorStatusCode(err)),
		Ip:            c.IP(),
		ContentLength: int64(c.Request().Header.ContentLength()),
	}
	log.Info().Interface("meta", meta).Msg(err.Error())

	errString, ok := err.(*pkgError.ErrorString)
	metaErrorCode := 500
	if ok {
		if errString.HttpCode() != 0 {
			metaErrorCode = errString.HttpCode()
		} else {
			metaErrorCode = errString.Code()
		}
	}
	return c.Status(metaErrorCode).JSON(response{
		Message: err.Error(),
		Data:    nil,
		Code:    getErrorStatusCode(err),
		Success: false,
	})
}
