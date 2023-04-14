package service

import (
	"fmt"
	"log"

	"github.com/begenov/tcp-chat/internal/pkg"
)

func (s *service) brodcaster() {
	for {
		s.mu.Lock()
		ch := <-s.c
		s.h.WriteString(ch.address + ch.text + "\n")
		for conn, name := range s.user {
			if name.name == ch.address {
				log.Printf("no")
				continue
			}

			if ch.text == "has left our chat..." || ch.text == "has joined our chat..." {
				_, err := fmt.Fprintln(conn, "\n"+ch.address+ch.text)
				if err != nil {
					continue
				}
			} else {

				_, err := fmt.Fprintln(conn, "\n"+ch.text)
				if err != nil {
					continue
				}
			}
			_, err := fmt.Fprintf(conn, "%s", fmt.Sprintf("[%s]:[%s]:", pkg.T(), name.name))
			if err != nil {
				continue
			}
		}

		s.mu.Unlock()
	}
}
