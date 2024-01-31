package goi18n

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestI18n(t *testing.T) {

	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("en.toml")
	bundle.MustLoadMessageFile("zh.toml")

	t.Logf("---- %v", bundle.LanguageTags())
	localizer := i18n.NewLocalizer(bundle)
	str, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    "0",
		TemplateData: nil,
	})
	t.Logf(" str: %v, err: %v", str, err)
}

func TestI18n2(t *testing.T) {
	engine, err := New("zh", "zh.toml")
	if err != nil {
		t.Logf("err: %v", err.Error())
		return
	}

	str1, err := engine.Get("0", nil)

	t.Logf("str1: %v, err: %v", str1, err)
}

func BenchmarkI18n(b *testing.B) {

	engine, err := New("zh", "zh.toml")
	if err != nil {
		b.Logf("err: %v", err.Error())
		return
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		engine.Get("0", nil)
	}
}
