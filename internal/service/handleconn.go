package service

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/begenov/tcp-chat/internal/pkg"
)

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
	username, err := s.username(scanner, conn)
	if err != nil {
		log.Printf("error in username: %v", err)
		return
	}
	if len(s.user) > 9 {
		_, err := fmt.Fprintln(conn, "Error Chat Full, come back in a while")
		if err != nil {
			log.Println("Couldn't fprint to the user")
			return
		}
		return
	}
	log.Println(username)
	s.user[conn] = &user{name: username}
	history, err := os.ReadFile(pathhistory)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(history) != 0 {
		_, err := fmt.Fprint(conn, string(history))
		if err != nil {
			log.Printf("handle conn in history: %v", err)
			return
		}
	}
	_, err = fmt.Fprintf(conn, "%s", fmt.Sprintf("[%s]:[%s]:", pkg.T(), username))
	if err != nil {
		log.Printf("error handle conn %v", err)
		return
	}

	join := message{text: " has joined our chat...", address: username}
	s.c <- join
	for scanner.Scan() {
		ln := pkg.Spacedeletion(scanner.Text())
		if ln == "" {
			log.Printf("error in space")
			continue
		}
		if !pkg.CheckSymbol(ln) {
			_, err := fmt.Fprintln(conn, "Error in text output\noutput text again")
			if err != nil {
				log.Printf("error")
				return
			}
		}
		_, err = fmt.Fprintf(conn, "%s", fmt.Sprintf("[%s]:[%s]:", pkg.T(), username))
		if err != nil {
			log.Printf("error %v", err)
			return
		}
		text := fmt.Sprintf("[%s]:[%s]:%s", pkg.T(), username, ln)
		msg := message{text: text, address: username}
		s.c <- msg
	}
	s.mu.Lock()
	delete(s.user, conn)
	s.mu.Unlock()
	left := message{text: " has left our chat...", address: username}
	s.c <- left

}
