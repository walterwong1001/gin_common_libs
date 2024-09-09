package validator

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/walterwong1001/gin_common_libs/enum/httpmethod"
)

var Translator ut.Translator

// 自定义方法验证器
func httpMethodValidator(fl validator.FieldLevel) bool {
	method := httpmethod.HTTPMethod(strings.ToUpper(fl.Field().String()))
	return method.IsValid()
}

func InitValidator() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//vl = validator.New()
		eng := en.New()
		uni := ut.New(eng, eng)
		Translator, _ = uni.GetTranslator("en")

		// 注册英语翻译
		if err := enTranslations.RegisterDefaultTranslations(v, Translator); err != nil {
			panic(fmt.Sprintf("Failed to register default translations: %v", err))
		}

		// 注册自定义验证器
		if err := v.RegisterValidation("http_method", httpMethodValidator); err != nil {
			panic(fmt.Sprintf("Failed to register validation: %v", err))
		}

		// 注册自定义翻译器
		//v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		//	name := fld.Tag.Get("json")
		//	if name == "" {
		//		name = fld.Name
		//	}
		//	return name
		//})

		// 注册自定义错误消息
		registerCustomTranslations(v, Translator)
	}
}

// 注册自定义错误消息
func registerCustomTranslations(validate *validator.Validate, trans ut.Translator) {
	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, translate)

	_ = validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} must be at least {1} characters", true)
	}, translate)

	_ = validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0} must be at most {1} characters", true)
	}, translate)

	_ = validate.RegisterTranslation("url", trans, func(ut ut.Translator) error {
		return ut.Add("url", "{0} must be a valid URL", true)
	}, translate)

	_ = validate.RegisterTranslation("http_method", trans, func(ut ut.Translator) error {
		return ut.Add("http_method", "{0} must be a valid HTTP method (GET, POST, PUT, DELETE, OPTIONS, HEAD)", true)
	}, translate)

	_ = validate.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} must be greater than or equal to {1}", true)
	}, translate)

	_ = validate.RegisterTranslation("lte", trans, func(ut ut.Translator) error {
		return ut.Add("lte", "{0} must be less than or equal to {1}", true)
	}, translate)

	_ = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email address", true)
	}, translate)
}

// 自定义翻译器
func translate(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T(fe.Tag(), fe.Tag(), fe.Param())
	return t
}
