package validator

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
)

type IValidator interface {
	SetValidator() IValidator

	GetValidator() *validator.Validate

	SetTranslator() IValidator

	GetTranslator() ut.Translator
}

type Validator struct {
	locale     string
	validator  *validator.Validate
	translator ut.Translator
}

func NewValidator(locale string) IValidator {
	return &Validator{locale: locale}
}

func (v *Validator) SetValidator() IValidator {
	v.validator = validator.New()

	return v
}

func (v *Validator) GetValidator() *validator.Validate {
	return v.validator
}

func (v *Validator) SetTranslator() IValidator {
	en := en.New()
	id := id.New()
	Uni := ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	err := enTranslations.RegisterDefaultTranslations(v.validator, transEN)
	if err != nil {
		fmt.Println(err)
	}

	err = idTranslations.RegisterDefaultTranslations(v.validator, transID)
	if err != nil {
		fmt.Println(err)
	}
	switch v.locale {
	case "en":
		v.translator = transEN
	case "id":
		v.translator = transID
	}

	return v
}

func (v *Validator) GetTranslator() ut.Translator {
	return v.translator
}
