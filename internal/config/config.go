package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

func init() {
	viper.AddConfigPath("$HOME/.dynamo")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO: parse and apply change
		log.Println("Config file changed:", e.Name)
	})
}
