package repository

import (
	"github.com/lekht/bookwiki-grpc/internal/models"
	"github.com/lekht/bookwiki-grpc/pkg/driver"
	"github.com/pkg/errors"
)

const (
	sqlBookByAuthor = `
	SELECT books.id, books.title
	FROM books 
	INNER JOIN author ON book.author_id = author.id 
	WHERE author.name = ?;`
)

type Repository struct {
	*driver.MySQL
}

func New(db *driver.MySQL) *Repository {
	return &Repository{db}
}

func (r *Repository) BooksByAuthor(author string) ([]models.Book, error) {
	var books []models.Book

	rows, err := r.DB.Query(sqlBookByAuthor, author)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make querry")
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Book

		if err := rows.Scan(&b.Id, &b.Title); err != nil {
			return nil, errors.Wrap(err, "failed to scan rows")
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to iterate rows")
	}

	return books, nil
}
