package configs

import (
	"github.com/alganbr/kedai-utils/configs"
	"github.com/spf13/viper"
	"strings"
)

func NewConfig() configs.Config {
	config, err := ReadInConfig()
	if err != nil {
		panic(err)
	}
	return *config
}

func ReadInConfig() (config *configs.Config, err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("env.yml")
	viper.AddConfigPath("./configs")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
