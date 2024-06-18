package config

import (
	"log"
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	configOnce sync.Once
	config     *AppConfig
)

type AppConfig struct {
	Server struct {
		Host string
		Port string
		Mode string
	}
	Log struct {
		Filename   string
		MaxSize    int `mapstructure:"max_size"`
		MaxBackups int `mapstructure:"max_backups"`
		MaxAge     int `mapstructure:"max_age"`
		Compress   bool
	}
	Db struct {
		Dsn string
	}
}

func GetConf() *AppConfig {
	configOnce.Do(func() {
		rootPath, _ := os.Getwd()
		viper.AddConfigPath(rootPath)
		viper.SetConfigName("config")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		initConfig := &AppConfig{}
		err = viper.Unmarshal(initConfig)
		if err != nil {
			panic(err)
		}
		config = initConfig
		log.Println("Config loaded")
		slog.Info("Config loaded")
	})
	return config
}
