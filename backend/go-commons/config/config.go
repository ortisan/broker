package config

import "github.com/spf13/viper"

type Config struct {
	Database      Database
	OpenTelemetry OpenTelemetry
	Server        Server
}

type Database struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

type OpenTelemetry struct {
	AgentHost string `mapstructure:"agent_host"`
	AgentPort string `mapstructure:"agent_port"`
}

type Server struct {
	Address string
	Name    string
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$PROJECT_HOME")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
