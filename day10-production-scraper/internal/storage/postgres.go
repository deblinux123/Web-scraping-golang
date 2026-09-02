package storage

import (
	"context"

	"github.com/deblinux123/Web-scraping-golang/day10-production-scraper/internal/model"
	"github.com/jackc/pgx/v5"
)

type PostgresStore struct {
	conn *pgx.Conn
}

func NewPostgresStore(conn *pgx.Conn) *PostgresStore {
	return &PostgresStore{
		conn: conn,
	}
}

func (s *PostgresStore) SaveBook(ctx context.Context, book model.Book) error {
	_, err := s.conn.Exec(
		ctx,
		`
		INSERT INTO books (
			title,
			price,
			availability,
			rating,
			url
		)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (url)
		DO UPDATE SET
			title = EXCLUDED.title,
			price = EXCLUDED.price,
			availability = EXCLUDED.availability,
			rating = EXCLUDED.rating,
			updated_at = NOW()
		`,
		book.Title,
		book.Price,
		book.Availability,
		book.Rating,
		book.URL,
	)

	return err
}
