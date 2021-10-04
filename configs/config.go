package configs

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var Config *viper.Viper

func init() {
	log.Info("Loading Db configs")

	configEnv := strings.ToLower(os.Getenv("ENV"))
	if len(configEnv) == 0 {
		configEnv = "dev"
	}
	Config = viper.New()
	Config.SetConfigName(configEnv)
	Config.AddConfigPath("./configs")
	Config.SetEnvPrefix("BACKEND")
	Config.SetConfigType("yaml")

	Config.AutomaticEnv()
	configLoadFile()
}

func configLoadFile() {
	// Load Configuration File
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatalf("ConfigLoadFile loading and reading error: %v\n", err)
	}
}
