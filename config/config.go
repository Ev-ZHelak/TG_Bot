package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Структура конфига
type TelegramConfig struct {
	Telegram struct {
		Token     string  `yaml:"token"`
		Debug     bool    `yaml:"debug"`
		Whitelist []int64 `yaml:"whitelist"`
	} `yaml:"telegram"`
}

type PostgreSQLConfig struct {
	Postgresql struct {
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		Username        string `yaml:"username"`
		Password        string `yaml:"password"`
		DBName          string `yaml:"dbName"`
		ApplicationName string `yaml:"applicationName"`
		SSLMode         string `yaml:"sslMode"`

		ConnectionPool struct {
			MaxConns    int           `yaml:"maxConns"`
			MinConns    int           `yaml:"minConns"`
			MaxLifetime time.Duration `yaml:"maxLifetime"`
			MaxIdleTime time.Duration `yaml:"maxIdleTime"`
		} `yaml:"connectionPool"`

		Timeouts struct {
			Connect time.Duration `yaml:"connect"`
			Query   time.Duration `yaml:"query"`
		} `yaml:"timeouts"`
	} `yaml:"postgresql"`
}

func LoadTelegramConfig() TelegramConfig {
	// Инициализация Viper
	viper.SetConfigFile("./config/telegram.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Ошибка загрузки yaml telegram:", err)
	}

	var cfg TelegramConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Ошибка парсинга конфига telegram:", err)
	}
	return cfg
}

func LoadPostgresConfig() PostgreSQLConfig {
	// Инициализация Viper
	viper.SetConfigFile("./config/postgres.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Ошибка загрузки yaml postgres:", err)
	}

	var cfg PostgreSQLConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Ошибка парсинга конфига postgres:", err)
	}
	return cfg
}
