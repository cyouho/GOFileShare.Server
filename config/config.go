package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
		// CertFile string `mapstructure:"cert_file"` 暂时不使用证书文件
		// KeyFile  string `mapstructure:"key_file"` 暂时不使用证书文件
	} `mapstructure:"server"`

	SharedDirectory string `mapstructure:"shared_directory"`
}

var Cfg *Config

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
}
