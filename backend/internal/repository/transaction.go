package repository

import (
	"github.com/AlexanderWangY/groupgo-backend/internal/db/sqlc"
	"github.com/jackc/pgx/v5"
)

func withOptionalTx(q *sqlc.Queries, txs ...pgx.Tx) *sqlc.Queries {
	if len(txs) > 0 {
		return q.WithTx(txs[0])
	}

	return q
}
