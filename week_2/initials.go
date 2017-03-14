package main

/*
Week 2 - Assignment 2

Design and implement a program, initials, that, given a person’s name, prints a person’s initials.
Implement your program in a file called initials.c in a directory called initials.
Your program should prompt a user for their name using get_string to obtain their name as a string.
You may assume that the user’s input will contain only letters (uppercase and/or lowercase) plus spaces. You don’t need to worry about names like Joseph Gordon-Levitt, Conan O’Brien, or David J. Malan!
But the user’s input might be sloppy, in which case there might be one or more spaces at the start and/or end of the user’s input or even multiple spaces in a row.
Your program should print the user’s initials (i.e., the first letter of each word in their name) with no spaces or periods, followed by a newline (\n).
*/

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
		// ASCII Ranges for alphabet
		lowercase := letter >= 97 && letter <= 122
		uppercase := letter >= 65 && letter <= 90
		space := letter == 32
		if space {
			seperator = true
			continue
		} else if seperator {
			if lowercase {
				//Didn't use helper function from strings
				letter -= 32
				fmt.Print(string(letter))
				seperator = false
			} else if uppercase {
				fmt.Print(string(letter))
				seperator = false
			}
		}
	}
}
