package pkg

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

func Spacedeletion(s string) string {
	var result string
	var count int
	for _, v := range s {
		if v != ' ' {
			if len(result) == 0 {
				count = 0
			}
			if count > 0 {
				result += " "
				count = 0
			}
			result += string(v)
		} else {
			count++
			continue
		}
	}
	return result
}

func CheckSymbol(s string) bool {
	s = strings.TrimSuffix(s, "\n")
	rxmsg := regexp.MustCompile("^[\u0400-\u04FF\u0020-\u007F]+$")
	if !rxmsg.MatchString(s) {
		return false
	}
	return true
}

func T() string {
	for {
		time := time.Now().Format("2006-01-02 15:04:05")
		return time
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
