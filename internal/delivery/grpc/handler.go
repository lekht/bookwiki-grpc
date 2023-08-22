package grpc

import (
	"context"

	"github.com/lekht/bookwiki-grpc/internal/delivery/grpc/wiki_grpc"
	"github.com/lekht/bookwiki-grpc/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	usecase usecase.WikiUsecase
	wiki_grpc.UnimplementedWikiHandlerServer
}

func NewWikiServer(gserver *grpc.Server, u *usecase.Usecase) {

	wikiServer := &server{
		usecase: u,
	}

	wiki_grpc.RegisterWikiHandlerServer(gserver, wikiServer)
	reflection.Register(gserver)
}

func (s *server) GetBooksByAuthor(ctx context.Context, r *wiki_grpc.AuthorRequest) (*wiki_grpc.BookListResponse, error) {
	author := r.GetAuthorName()

	books, err := s.usecase.GetBooksByAuthor(author)
	if err != nil {
		return nil, err
	}

	var booksList []*wiki_grpc.Book
	for _, b := range books {
		booksList = append(booksList, &wiki_grpc.Book{
			Id:    int32(b.Id),
			Title: b.Title,
		})
	}

	return &wiki_grpc.BookListResponse{
		Books: booksList,
	}, nil
}

func (s *server) GetAuthorsByBook(ctx context.Context, r *wiki_grpc.BookRequest) (*wiki_grpc.AuthorListResponse, error) {
	book := r.GetBookTitle()

	authors, err := s.usecase.GetAuthorsByBook(book)
	if err != nil {
		return nil, err
	}

	var authorsList []*wiki_grpc.Author
	for _, a := range authors {
		authorsList = append(authorsList, &wiki_grpc.Author{
			Id:         int32(a.Id),
			AuthorName: a.AuthorName,
		})
	}

	return &wiki_grpc.AuthorListResponse{
		Authors: authorsList,
	}, nil
}
