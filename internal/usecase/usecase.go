package usecase

import (
	"log"

	"github.com/lekht/bookwiki-grpc/internal/models"
	"github.com/lekht/bookwiki-grpc/internal/repository"
	"github.com/pkg/errors"
)

var (
	errNotFound       = errors.New("Your requested Item is not found")
	errServerInternal = errors.New("server internal error")
)

type WikiUsecase interface {
	GetAuthorsByBook(book string) ([]models.Author, error)
	GetBooksByAuthor(author string) ([]models.Book, error)
}

type Usecase struct {
	RepoUsecase
}

type RepoUsecase interface {
	AuthorsByBook(book string) ([]models.Author, error)
	BooksByAuthor(author string) ([]models.Book, error)
}

func New(r *repository.Repository) *Usecase {
	return &Usecase{r}
}

func (u *Usecase) GetAuthorsByBook(book string) ([]models.Author, error) {

	authors, err := u.AuthorsByBook(book)
	if err != nil {
		log.Println("failed to make repository request: ", err)
		return nil, errServerInternal
	}

	if authors == nil {
		return nil, errNotFound
	}

	return authors, nil
}

func (u *Usecase) GetBooksByAuthor(author string) ([]models.Book, error) {

	books, err := u.BooksByAuthor(author)
	if err != nil {
		log.Println("failed to make repository request: ", err)
		return nil, errServerInternal
	}

	if books == nil {
		return nil, errNotFound
	}

	return books, nil
}
