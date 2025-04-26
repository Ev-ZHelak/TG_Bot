package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Структура для конфига
type ConfigTgBot struct {
	Telegram struct {
		Token     string  `yaml:"token"`
		Debug     bool    `yaml:"debug"`
		Whitelist []int64 `yaml:"whitelist"`
	} `yaml:"telegram"`
}

type ConfigSgl struct {
	Database struct {
		URL            string        `yaml:"url"`
		Timeout        time.Duration `yaml:"timeout"`
		MaxConnections int           `yaml:"maxConnections"`
	} `yaml:"database"`
}

func LoadTgBotConfig() ConfigTgBot {
	// Инициализация Viper
	viper.SetConfigFile("./internal/config/app.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Ошибка загрузки yaml:", err)
	}

	var cfg ConfigTgBot

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Ошибка парсинга конфига:", err)
	}
	return cfg
}

func LoadSqlConfig() ConfigSgl {
	// Инициализация Viper
	viper.SetConfigFile("./internal/config/app.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Ошибка загрузки yaml:", err)
	}

	var cfg ConfigSgl

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Ошибка парсинга конфига:", err)
	}
	return cfg
}
