package service

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type Service interface {
	Run(string) error
}

type service struct {
	user map[net.Conn]*user
	mu   sync.Mutex
	c    chan message
	h    os.File
}

const pathwelcome = "welcome.txt"
const pathhistory = "history.txt"

func NewService() Service {
	return &service{
		user: make(map[net.Conn]*user),
		c:    make(chan message),
		h:    os.File{},
	}
}

func (s *service) Run(port string) error {
	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("error net.Listen: %w", err)
		return err
	}
	defer li.Close()
	h, err := os.Create(pathhistory)
	if err != nil {
		fmt.Println(err)
		return err
	}
	s.h = *h
	go s.brodcaster()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("error Accept: %w", err)
			continue
		}
		go s.handleConn(conn)
	}
}
