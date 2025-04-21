package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Структура для конфига
type Config struct {
	Telegram struct {
		Token     string  `yaml:"token"`
		Debug     bool    `yaml:"debug"`
		Whitelist []int64 `yaml:"whitelist"`
	} `yaml:"telegram"`
	Database struct {
		URL            string        `yaml:"url"`
		Timeout        time.Duration `yaml:"timeout"`
		MaxConnections int           `yaml:"max_connections"`
	} `yaml:"database"`
}

func LoadTgBotConfig() Config {
	// Инициализация Viper
	viper.SetConfigFile("./internal/config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Ошибка загрузки yaml:", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Ошибка парсинга конфига:", err)
	}

	fmt.Printf("Token: %s\n", cfg.Telegram.Token)
	fmt.Printf("Timeout db: %v\n", cfg.Database.Timeout)
	return cfg
}
