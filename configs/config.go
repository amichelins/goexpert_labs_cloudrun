package configs

import (
    "errors"
    "strings"

    "github.com/spf13/viper"
)

type conf struct {
    WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
    WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

var ErrorMissingKey = errors.New("WeatherApiKey não configurada")

func LoadConfig(path string) (*conf, error) {
    var cfg *conf
    viper.SetConfigName("app_config")
    viper.SetConfigType("env")
    viper.AddConfigPath(path)
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()

    if err != nil {
        panic(err)
    }

    err = viper.Unmarshal(&cfg)

    if err != nil {
        panic(err)
    }

    if strings.TrimSpace(cfg.WeatherApiKey) == "" {
        return nil, ErrorMissingKey
    }

    if strings.TrimSpace(cfg.WebServerPort) == "" {
        cfg.WebServerPort = ":8000"
    }

    return cfg, err
}
