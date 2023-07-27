package config

import "github.com/spf13/viper"

type Config struct {

	//mysql setup
	DBHost     string `mapstructure:"SQL_HOST"`
	DBUsername string `mapstructure:"SQL_USER"`
	DBPassword string `mapstructure:"SQL_PASS"`
	DBName     string `mapstructure:"SQL_DB_NAME"`
	DBPort     string `mapstructure:"SQL_PORT"`

	// redis setup
	RedisUrl string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// handler null
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	_ = viper.Unmarshal(&config)

	return
}
