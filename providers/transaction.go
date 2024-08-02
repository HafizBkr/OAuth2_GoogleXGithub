package providers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransactionProvider struct {
	db *sqlx.DB
}

func NewTransactionProvider(db *sqlx.DB) *TransactionProvider {
	return &TransactionProvider{
		db: db,
	}
}

func (p *TransactionProvider) Provide() (*sqlx.Tx, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("Error while providing transaction: %w", err)
	}
	return tx, nil
}
