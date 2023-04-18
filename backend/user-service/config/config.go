package config

import "github.com/spf13/viper"

type config struct {
	Database struct {
		User         string
		Password     string
		Host         string
		Port         string
		DatabaseName string
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
