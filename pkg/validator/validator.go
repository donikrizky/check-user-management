package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	pkgError "gitlab.playcourt.id/telkom-digital/svc/spbe/perizinan-event-user-management/pkg/error"
)

type XValidator struct {
	Validator *validator.Validate
}

var validate = validator.New()

func (v XValidator) Validate(data interface{}) error {
	var message string
	if errs := validate.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			message = message + strings.ToLower(fmt.Sprintf("%s %s;", err.Field(), err.Tag()))
		}
		return pkgError.BadRequest(message)
	}

	return nil
}
