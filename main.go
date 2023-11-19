package main

import (
	"log"
)

func main() {
  storage, err := NewPostgresStore()
  if err != nil {
    log.Fatal(err)
  }

  if err := storage.Init(); err != nil {
    log.Fatal(err)
  }

  server := NewAPIServer(":3000", storage)
  server.Run()
}
