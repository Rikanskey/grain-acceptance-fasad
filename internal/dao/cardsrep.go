package dao

import "database/sql"

type CardRepository struct {
	cards *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{cards: db}
}
