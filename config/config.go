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

	WindowsSharedDirectory string `mapstructure:"windows_shared_directory"`
	LinuxSharedDirectory   string `mapstructure:"linux_shared_directory"`
}

var Cfg *Config

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
}
