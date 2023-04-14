package service

import (
	"bufio"
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
}

type user struct {
	name string
}

type message struct {
	text    string
	address string
}

const pathwelcome = "welcome.txt"

func NewService() Service {
	return &service{
		user: make(map[net.Conn]*user),
		c:    make(chan message),
	}
}

func (s *service) Run(port string) error {
	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("error net.Listen: %w", err)
		return err
	}
	defer li.Close()

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

func (s *service) handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	file, err := os.ReadFile(pathwelcome)
	if err != nil {
		log.Printf("error read file in welcome.txt: %v", err)
		return
	}
	_, err = fmt.Fprint(conn, string(file))
	if err != nil {
		log.Println("Couldn't fprint to the user")
		return
	}
	log.Println(scanner.Text())
}

func (s *service) brodcaster() {
}
