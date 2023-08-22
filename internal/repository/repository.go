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
	INNER JOIN authors ON books.author_id = authors.id 
	WHERE authors.author_name = ?;`

	sqlAuthorByBook = `
	SELECT authors.id, authors.author_name 
	FROM authors
	INNER JOIN books ON authors.id = books.author_id
	WHERE books.title = ?;
`
)

type Repository struct {
	*driver.MySQL
}

func New(db *driver.MySQL) *Repository {
	return &Repository{db}
}

func (r *Repository) AuthorsByBook(book string) ([]models.Author, error) {
	var authors []models.Author

	rows, err := r.DB.Query(sqlAuthorByBook, book)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make querry")
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Author

		if err := rows.Scan(&a.Id, &a.AuthorName); err != nil {
			return nil, errors.Wrap(err, "failed to scan rows")
		}
		authors = append(authors, a)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to iterate rows")
	}

	return authors, nil
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
