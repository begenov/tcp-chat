package main

import (
	"log"
	"os"

	"github.com/begenov/tcp-chat/internal/pkg"
	"github.com/begenov/tcp-chat/internal/service"
)

var port = ":8989"

func main() {
	log.Println("Starting tcp-chat ...")
	var err error
	if len(os.Args[1:]) > 1 || len(os.Args[1:]) != 0 {
		log.Printf("[USAGE]: ./TCPChat %v", port)
		return
	}
	if len(os.Args[1:]) == 1 {
		port, err = pkg.CheckPort(os.Args[1:][0])
		if err != nil {
			log.Printf("Check port error: %v", err)
		}
	}

	log.Printf("Listening on the port: %v", port)
	s := service.NewService()
	if err := s.Run(port); err != nil {
		log.Printf("Error in Run function: %v", err)
	}
}
