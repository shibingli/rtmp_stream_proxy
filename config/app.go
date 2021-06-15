package config

type App struct {
	AppName  string `json:"app_name" yaml:"app_name" mapstructure:"app_name"`
	LogLevel string `json:"log_level" yaml:"log_level" mapstructure:"log_level"`
}
