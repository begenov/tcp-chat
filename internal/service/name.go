package service

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/begenov/tcp-chat/internal/pkg"
)

func (s *service) username(scanner *bufio.Scanner, conn net.Conn) (string, error) {
	var username string
	for {
		_, err := fmt.Fprint(conn, "[ENTER YOUR NAME]:")
		if err != nil {
			log.Printf("error username: %v", err)
			return "", err
		}
		if scanner.Scan() {
			username = pkg.Spacedeletion(scanner.Text())
			if !pkg.CheckSymbol(scanner.Text()) {
				_, err := fmt.Fprint(conn, "Error in text output\noutput text again\n")
				if err != nil {
					return "", err
				}
				continue
			}
			if username == "" {
				_, err := fmt.Fprintln(conn, "Error the void is not allowed")
				if err != nil {
					return "", err
				}
				continue
			}
			if !s.repeatName(pkg.Spacedeletion(scanner.Text())) {
				_, err := fmt.Fprintln(conn, "Error repeat in name")
				if err != nil {
					return "", err
				}
				continue
			}

			break
		}
	}
	return username, nil
}

func (s *service) repeatName(user string) bool {
	for _, name := range s.user {
		if name.name == user {
			return false
		}
	}
	return true
}
