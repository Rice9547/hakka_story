package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Auth0    Auth0Config    `mapstructure:"auth0"`
	Upload   SpaceConfig    `mapstructure:"upload"`
}

type ServerConfig struct {
	URL  string `mapstructure:"url"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type Auth0Config struct {
	Domain   string `mapstructure:"domain"`
	Audience string `mapstructure:"audience"`
}

type SpaceConfig struct {
	Endpoint    string `mapstructure:"endpoint"`
	ImageBucket string `mapstructure:"image_bucket"`
	AudioBucket string `mapstructure:"audio_bucket"`
	Region      string `mapstructure:"region"`
	AccessKey   string `mapstructure:"access_key"`
	SecretKey   string `mapstructure:"secret_key"`
}

func LoadConfig(configPath string) (Config, error) {
	var config Config

	viper.SetConfigFile(configPath)
	viper.SetEnvPrefix("hakka")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return config, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return config, err
	}

	return config, nil
}
