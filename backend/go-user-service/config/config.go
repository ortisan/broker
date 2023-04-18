package config

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		Host         string
		Port         string
		User         string
		Password     string
		DatabaseName string
	}

	OpenTelemetry struct {
		AgentHost string `mapstructure:"agent_host"`
		AgentPort string `mapstructure:"agent_port"`
	}

	Server struct {
		Address string
		Name    string
	}
}

func NewConfig() (*Config, error) {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$BROKER_PROJECT_HOME/go-user-service")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
