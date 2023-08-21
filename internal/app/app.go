package app

import (
	"log"

	"github.com/lekht/bookwiki-grpc/pkg/driver"
)

func Run() {
	db, err := driver.New()
	if err != nil {
		log.Panicf("failed to create new driver connection: %v", err)
	}

	defer db.Close()

}
