package db

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PrintStatistics(p *pgxpool.Pool) {
	stats := p.Stat()
	fmt.Printf(
		"Соединения: %d (активные: %d, простаивающие: %d)\n",
		stats.TotalConns(),
		stats.AcquiredConns(),
		stats.IdleConns(),
	)
}
