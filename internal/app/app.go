package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	deliveryGrpc "github.com/lekht/bookwiki-grpc/internal/delivery/grpc"
	"github.com/lekht/bookwiki-grpc/internal/repository"
	"github.com/lekht/bookwiki-grpc/internal/usecase"
	"github.com/lekht/bookwiki-grpc/pkg/driver"
	"github.com/lekht/bookwiki-grpc/pkg/grpcserver"
	"google.golang.org/grpc"
)

func Run() {
	db, err := driver.New()
	if err != nil {
		log.Panicf("failed to create new driver connection: %s\n", err)
	}

	repo := repository.New(db)

	defer db.Close()

	u := usecase.New(repo)

	server := grpc.NewServer()

	deliveryGrpc.NewWikiServer(server, u)

	gServer := grpcserver.New(server, os.Getenv("SERVER_ADDRESS"))
	log.Println("Server Run at ", os.Getenv("SERVER_ADDRESS"))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("interrupt sinal: " + s.String())
	case err = <-gServer.Notify():
		log.Println(fmt.Errorf("grpc server error: %w", err))
	}

	gServer.Stop()
}
