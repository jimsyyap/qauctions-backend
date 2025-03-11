package config

import "github.com/spf13/viper"

type Config struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string
}

func LoadConfig() *Config {
    viper.SetDefault("DB_HOST", "localhost")
    viper.SetDefault("DB_USER", "postgres")
    viper.SetDefault("DB_PASSWORD", "password")
    viper.SetDefault("DB_NAME", "auctions")
    viper.SetDefault("DB_PORT", "5432")

    return &Config{
        DBHost:     viper.GetString("DB_HOST"),
        DBUser:     viper.GetString("DB_USER"),
        DBPassword: viper.GetString("DB_PASSWORD"),
        DBName:     viper.GetString("DB_NAME"),
        DBPort:     viper.GetString("DB_PORT"),
    }
}
