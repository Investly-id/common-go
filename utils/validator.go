package utils

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

type FieldError struct {
	Param   string
	Message string
}

func msgForTag(fe validator.FieldError, fieldName string) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("Kolom %s wajib diisi", fieldName)
	case "numeric":
		return fmt.Sprintf("Kolom %s hanya boleh angka", fieldName)
	case "email":
		return fmt.Sprintf("Email tidak valid")
	}
	return fe.Error()
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		if name == "" {
			name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		}

		return name
	})
	if err := cv.Validator.Struct(i); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorRes := make(map[string]interface{})
			for _, fe := range ve {
				errorRes[fe.Field()] = msgForTag(fe, fe.Field())
			}
			return echo.NewHTTPError(http.StatusBadRequest, errorRes)
		}
	}
	return nil
}

const (
	MB             = 1 << 20
	MAX_IMAGE_SIZE = 3 * MB
)

func FileValidation(c echo.Context, fieldName string, isRequired bool, fileType string) error {
	form, _ := c.MultipartForm()

	files := form.File[fieldName]
	if files == nil && isRequired {
		return errors.New(fmt.Sprintf("Kolom %s wajib diisi", fieldName))
	}

	if files != nil {
		for _, i := range files {

			ext := strings.ToLower(path.Ext(i.Filename))

			if fileType == "image" {
				if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
					return errors.New("Hanya boleh gambar dengan format (.jpg, .jpeg, .png)")
				}
			}

			if i.Size > MAX_IMAGE_SIZE {
				return errors.New(fmt.Sprintf("Maksimal ukuran gambar = %d", MAX_IMAGE_SIZE))
			}
		}
	}

	return nil
}

func GetErrorValidation(err error) interface{} {
	return err.(*echo.HTTPError).Message
}
