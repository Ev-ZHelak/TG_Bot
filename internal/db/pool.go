package db

import (
	"TG_Bot/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPostgresPool(config config.PostgreSQLConfig) (*pgxpool.Pool, error) {
	if Pool != nil {
		return Pool, nil
	}
	// Формируем строку подключения
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s application_name=%s sslmode=%s",
		config.Postgresql.Host,
		config.Postgresql.Port,
		config.Postgresql.Username,
		config.Postgresql.Password,
		config.Postgresql.DBName,
		config.Postgresql.ApplicationName,
		config.Postgresql.SSLMode,
	)

	// Создаем конфигурацию пула
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("не удалось проанализировать строку подключения: %w", err)
	}

	// Настраиваем параметры пула соединений
	poolConfig.MaxConns = int32(config.Postgresql.ConnectionPool.MaxConns)
	poolConfig.MinConns = int32(config.Postgresql.ConnectionPool.MinConns)
	poolConfig.MaxConnLifetime = config.Postgresql.ConnectionPool.MaxLifetime
	poolConfig.MaxConnIdleTime = config.Postgresql.ConnectionPool.MaxIdleTime
	poolConfig.HealthCheckPeriod = config.Postgresql.ConnectionPool.HealthCheckPeriod

	// Настраиваем таймауты
	poolConfig.ConnConfig.ConnectTimeout = config.Postgresql.Timeouts.Connect
	poolConfig.ConnConfig.RuntimeParams["statement_timeout"] = fmt.Sprintf("%d", config.Postgresql.Timeouts.Query.Milliseconds())

	// Создаем пул соединений с контекстом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пул соединений: %w", err)
	}

	// Проверяем подключение
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("не удалось пинговать базу данных: %w", err)
	}
	return pool, nil
}
