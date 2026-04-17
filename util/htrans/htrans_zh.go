package htrans

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 初始化
func init() {
	InitTrans()
}

var Uni *ut.UniversalTranslator

// InitTrans 初始化翻译器，修改gin框架中的 validator 引擎, 注册支持的翻译器
func InitTrans() error {
	// 修改gin框架中的 validator 引擎属性, 实现自定制
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法，返回错误字段使用 json tag 字段，而不是结构体字段名
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		enT := en.New()             // 英文翻译器
		zhT := zh.New()             // 中文翻译器
		Uni = ut.New(enT, zhT, enT) // 第一个参数是备用语言环境，后面的参数是应该支持的语言环境

		// 注册中文翻译器
		if trans, ok := Uni.GetTranslator("zh"); ok {
			if err := zhTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
				return err
			}
		} else {
			return fmt.Errorf(`Uni.GetTranslator("%s") error`, "zh")
		}

		// 注册英文翻译器
		if trans, ok := Uni.GetTranslator("en"); ok {
			if err := enTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
				return err
			}
		} else {
			return fmt.Errorf(`Uni.GetTranslator("%s") error`, "en")
		}
		return nil
	} else {
		return fmt.Errorf("")
	}
}

// GetLocalTrans 根据入参语言类型，注册相应的翻译器
// 项目中通常会将该函数放到中间件中, locale 会从请求的 headers 中获取
func GetLocalTrans(locale string) (ut.Translator, error) {
	// 根据不同语言获取相应的翻译器
	if trans, ok := Uni.GetTranslator(locale); !ok {
		return nil, fmt.Errorf(`Uni.GetTranslator("%s") error`, locale)
	} else {
		return trans, nil
	}
}

/*
removeTopStruct 移除验证错误字段最前名的结构体名称
移除前

	{
	    "error": {
	        "Family.member.email": "email必须是一个有效的邮箱",
	        "Family.member.name": "name长度必须至少为3个字符",
	        "Family.name": "name长度必须至少为3个字符"
	    }
	}

移除后

	{
	    "error": {
	        "member.email": "email必须是一个有效的邮箱",
	        "member.name": "name长度必须至少为3个字符",
	        "name": "name长度必须至少为3个字符"
	    }
	}
*/
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// 翻译
func Trans(err error) (res any) {
	tarns, _ := GetLocalTrans("zh") // 翻译中文
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		res = err.Error()
	} else {
		res = RemoveTopStruct(errs.Translate(tarns))
	}
	return
}
