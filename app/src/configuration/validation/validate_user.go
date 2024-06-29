package validation

import (
	"encoding/json"
	"errors"
	"golang-crud/app/src/configuration/app_errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {

	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}

}

func ValidateUserError(validation_err error) *app_errors.AppError {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	var finalError *app_errors.AppError

	if errors.As(validation_err, &jsonErr) {
		finalError = app_errors.NewBadRequestError("Invalid field type", nil)

	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []app_errors.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := app_errors.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		finalError = app_errors.NewBadRequestError("Some fields are invalid", errorsCauses)

	} else {
		finalError = app_errors.NewBadRequestError("Error trying to convert fields", nil)
	}

	return finalError

}
