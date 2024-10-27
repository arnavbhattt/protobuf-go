package main

import (
	"log"

	"github.com/arnavbhattt/protobuf-go/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
