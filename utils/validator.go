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
	MAX_PDF_SIZE   = 2 * MB
	FILE_IMAGE     = "image"
	FILE_PDF       = "pdf"
)

// Deprecated: Use FileValidationV2
func FileValidation(c echo.Context, fieldName string, isRequired bool, fileType string) error {
	form, _ := c.MultipartForm()

	files := form.File[fieldName]
	if files == nil && isRequired {
		return errors.New(fmt.Sprintf("Kolom %s wajib diisi", fieldName))
	}

	if files != nil {
		for _, i := range files {

			ext := strings.ToLower(path.Ext(i.Filename))

			if fileType == FILE_IMAGE {
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

func FileValidationV2(c echo.Context, fieldName string, isRequired bool, fileTypes []string) error {
	form, _ := c.MultipartForm()

	files := form.File[fieldName]
	if files == nil && isRequired {
		return errors.New(fmt.Sprintf("Kolom %s wajib diisi", fieldName))
	}

	imageType := []string{".jpg", ".jpeg", ".png"}
	pdfType := []string{".pdf"}

	// Define allowed formats
	allowedFormats := make([]string, 0)
	for _, fileType := range fileTypes {
		if fileType == FILE_IMAGE {
			allowedFormats = append(allowedFormats, imageType...)
		} else if fileType == FILE_PDF {
			allowedFormats = append(allowedFormats, pdfType...)
		}
	}

	if files != nil {
		for _, i := range files {

			ext := strings.ToLower(path.Ext(i.Filename))

			flag := false
			for _, allowedFormat := range allowedFormats {
				if ext == allowedFormat {
					flag = true
				}
			}

			// return if ext doesn't exist in allowed format
			if !flag {
				return errors.New(fmt.Sprintf("File hanya boleh dengan format (%s)", strings.Join(allowedFormats, ", ")))
			}

			// Check image size
			for _, fileTypeImage := range imageType {
				if ext == fileTypeImage {
					if i.Size > MAX_IMAGE_SIZE {
						return errors.New(fmt.Sprintf("Maksimal ukuran gambar = %d", MAX_IMAGE_SIZE))
					}
				}
			}

			// Check pdf size
			for _, fileTypePdf := range pdfType {
				if ext == fileTypePdf {
					if i.Size > MAX_PDF_SIZE {
						return errors.New(fmt.Sprintf("Maksimal ukuran dokumen = %d", MAX_PDF_SIZE))
					}
				}
			}
		}
	}

	return nil
}

func GetErrorValidation(err error) interface{} {
	return err.(*echo.HTTPError).Message
}
