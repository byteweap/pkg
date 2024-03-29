package goi18n

import (
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/cast"
	"golang.org/x/text/language"
)

type Engine struct {
	loader *i18n.Localizer
}

func New(lang string, path string) (*Engine, error) {

	lan := language.Make(lang)
	if lan.IsRoot() {
		return nil, errors.New("language is not supported")
	}

	bundle := i18n.NewBundle(lan)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(path)
	eg := &Engine{
		loader: i18n.NewLocalizer(bundle, lan.String()),
	}
	return eg, nil
}

// Get 根据(id)key获取value
// 支持模板数据填充，非模版数据 templateData为nil即可
func (e *Engine) Get(id interface{}, templateData map[string]interface{}) (string, error) {

	idStr := cast.ToString(id)
	if idStr == "" {
		return "", errors.New("id is empty")
	}
	return e.loader.Localize(&i18n.LocalizeConfig{MessageID: idStr, TemplateData: templateData})
}
