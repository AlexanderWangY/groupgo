package transaction

import (
	"context"
	"fmt"

	"github.com/AlexanderWangY/groupgo-backend/internal/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txFunc func(q *sqlc.Queries, tx pgx.Tx) error

type TransactionManager struct {
	db *pgxpool.Pool
}

func NewTransactionManager(db *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{
		db: db,
	}
}

func (t *TransactionManager) WithTx(ctx context.Context, fn txFunc) error {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("[WithTx] error starting transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		}
	}()

	q := sqlc.New(tx)

	if err := fn(q, tx); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("[WithTx] rollback error: %v (original error: %w)", rbErr, err)
		}

		return fmt.Errorf("[WithTx] transaction error: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("[WithTx] commit error: %w", err)
	}

	return nil
}
