package config

import (
	"log"
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
	})
	return config
}
