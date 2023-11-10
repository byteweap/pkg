package configx

import (
	"github.com/byteweap/pkg/logs"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// LoadConfigFromFile loads the configuration from a file.
//
// It takes the following parameters:
// - filename (string): the name of the file to load the configuration from.
// - v (any): the configuration structure to populate with the loaded data.
// - listen (bool): a flag indicating whether to listen for configuration changes.
// It returns an error if there was a problem loading the configuration.
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

// load loads a configuration file into a Viper instance and unmarshals it into a provided struct.
//
// vp: A pointer to a Viper instance.
// filename: The path to the configuration file.
// v: The struct to unmarshal the configuration into.
// error: Returns an error if the configuration file cannot be read or if unmarshaling fails.
func load(vp *viper.Viper, filename string, v any) error {

	vp.SetConfigFile(filename)
	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	return vp.Unmarshal(v)
}
