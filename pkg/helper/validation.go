package helper

import (
	"net/mail"
	"strings"

	"github.com/tuongnguyen1209/poly-career-back/pkg/response"
	validationMessage "github.com/tuongnguyen1209/poly-career-back/pkg/validation"
	"golang.org/x/exp/slices"
)

type Validation struct {
	IsValid    bool
	ErrorField []response.ErrorField
}

func InitValidation() *Validation {
	return &Validation{
		IsValid:    true,
		ErrorField: []response.ErrorField{},
	}
}

func getField(errorField []response.ErrorField, field string) int {
	return slices.IndexFunc(errorField, func(ef response.ErrorField) bool {
		return ef.Field == field
	})
}

func (v *Validation) SetError(field, code string) *Validation {
	code = fieldMessage(code, field)
	inx := getField(v.ErrorField, field)
	v.IsValid = false
	if inx >= 0 {
		v.ErrorField[inx].Code = slices.Insert(v.ErrorField[inx].Code, len(v.ErrorField[inx].Code), code)
	} else {
		v.ErrorField = slices.Insert(v.ErrorField, len(v.ErrorField), response.ErrorField{
			Field: field,
			Code:  []string{code},
		})
	}
	return v
}

func (v *Validation) Require(field string, value interface{}) *Validation {
	checked := false
	switch value.(type) {
	case string:
		checked = value == "" || value == nil
	default:
		checked = value == nil
	}

	if !checked {
		return v
	}

	return v.SetError(field, validationMessage.Missing)
}

func (v *Validation) IsEmail(field, value string) *Validation {
	if _, err := mail.ParseAddress(value); err != nil {
		return v.SetError(field, validationMessage.IsValid)
	}

	return v
}

func fieldMessage(text, field string) string {
	newText := strings.ReplaceAll(text, "<field>", field)
	return newText
}
