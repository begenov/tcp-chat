package main

import (
	"fmt"
	"log"
	"os"

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
		port, err = CheckPort(os.Args[1:][0])
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

func CheckPort(s string) (string, error) {
	var result string
	if len(s) != 4 {
		log.Println("the port check for the length according to the diphtholt it should be 4")
		return "", fmt.Errorf("the port check for the length")
	}
	for _, v := range s {
		if v < '0' || v > '9' {
			log.Println("There must be numbers")
			return "", fmt.Errorf("there must be numbers")
		} else {
			result += string(v)
		}
	}
	return ":" + result, nil
}
