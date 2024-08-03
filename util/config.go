package util

import "github.com/spf13/viper"

type Config struct {
	Port uint16 `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

func LoadConfig(path string) (config *Config, err error ) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return  nil, err
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, err
	}

	return config, nil
}