package configx

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/violin8/pkg/logs"
)

// LoadConfigFromFile 加载配置文件,并解析到结构体
// filename: 完整文件路径 如: config/cfg.toml
// v: 结构体指针
// listen: 是否监听文件变化
func LoadConfigFromFile(filename string, v any, listen bool) error {
	vp := viper.New()
	if err := load(vp, filename, v); err != nil {
		return err
	}

	if listen {
		vp.WatchConfig()
		vp.OnConfigChange(func(_ fsnotify.Event) {
			if err := load(vp, filename, v); err != nil {
				logs.Errorx().Any("Filename", filename).Any("Error", err.Error()).Msg("--------OnConfigChange")
			} else {
				logs.Infox().Any("Filename", filename).Any("Data", v).Msg("--------OnConfigChange")
			}
		})
	}

	return nil
}

// 加载,解析
func load(vp *viper.Viper, filename string, v any) error {

	vp.SetConfigFile(filename)
	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	return vp.Unmarshal(v)
}
