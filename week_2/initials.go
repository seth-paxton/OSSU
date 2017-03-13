package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Name: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		initials(text)
		break
	}
}

func initials(s string) {
	seperator := true
	for _, letter := range s {
		lowercase := letter >= 97 && letter <= 122
		uppercase := letter >= 65 && letter <= 90
		space := letter == 32
		if space {
			seperator = true
			continue
		} else if lowercase || uppercase {
			if seperator {
				if lowercase {
					letter -= 32
					fmt.Print(string(letter))
					seperator = false
				} else {
					fmt.Print(string(letter))
					seperator = false
				}
			}
		}
	}
}
