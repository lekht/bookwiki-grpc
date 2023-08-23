package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	pb "github.com/lekht/bookwiki-grpc/internal/delivery/grpc/wiki_grpc"
	"github.com/lekht/bookwiki-grpc/internal/repository"
	"github.com/lekht/bookwiki-grpc/internal/usecase"
	"github.com/lekht/bookwiki-grpc/pkg/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var sqlMock sqlmock.Sqlmock

func init() {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection\n", err)
	}
	myDB := driver.MySQL{DB: db}
	sqlMock = mock

	r := repository.New(&myDB)
	uc := usecase.New(r)

	s := grpc.NewServer()

	NewWikiServer(s, uc)

	lis = bufconn.Listen(bufSize)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Test_server_GetAuthorsByBook(t *testing.T) {
	querry := "SELECT authors.id, authors.author_name FROM authors INNER JOIN books ON authors.id = books.author_id WHERE books.title = \\?;"
	rows := sqlmock.NewRows([]string{"id", "author_name"}).AddRow(1, "author_test")
	sqlMock.ExpectQuery(querry).WithArgs("book_test").WillReturnRows(rows)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewWikiHandlerClient(conn)

	resp, err := client.GetAuthorsByBook(ctx, &pb.BookRequest{BookTitle: "book_test"})
	if err != nil {
		t.Fatalf("GetAuthorsByBook failed: %v", err)
	}
	fmt.Printf("Response: %+v", resp)
}

func Test_server_GetBooksByAuthor(t *testing.T) {
	querry := "SELECT books.id, books.title	FROM books INNER JOIN authors ON books.author_id = authors.id WHERE authors.author_name = \\?;"
	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "testbook1").
		AddRow(2, "testbook2")

	sqlMock.ExpectQuery(querry).WithArgs("testAuthor1").WillReturnRows(rows)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewWikiHandlerClient(conn)

	resp, err := client.GetBooksByAuthor(ctx, &pb.AuthorRequest{AuthorName: "testAuthor1"})
	if err != nil {
		t.Fatalf("GetAuthorsByBook failed: %v", err)
	}
	fmt.Printf("Response: %+v", resp)
}
