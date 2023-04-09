package helper

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisTTL      string `mapstructure:"REDIS_TTL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
