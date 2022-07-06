package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var Trans ut.Translator

func InitTranslate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)

		Trans, _ = uni.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(v, Trans)
	}
}