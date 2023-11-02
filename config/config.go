package config

import "github.com/spf13/viper"

func Load() (Cfg, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	var cfg Cfg
	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

type Cfg struct {
	WebServerPort  string `mapstructure:"WEB_SERVER_PORT"`
	GrpcServerPort string `mapstructure:"GRPC_SERVER_PORT"`
}
