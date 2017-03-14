package main

/*
Week 2 - Assignment 2

Design and implement a program, caesar, that encrypts messages using Caesar’s cipher.
Implement your program in a file called caesar.c in a directory called caesar.
Your program must accept a single command-line argument, a non-negative integer. Let’s call it k for the sake of discussion.
If your program is executed without any command-line arguments or with more than one command-line argument, your program should print an error message of your
choice (with printf) and return from main a value of 1 (which tends to signify an error) immediately.
You can assume that, if a user does provide a command-line argument, it will be a non-negative integer (e.g., 1). No need to check that it’s indeed numeric.
Do not assume that k will be less than or equal to 26. Your program should work for all non-negative integral values of k less than 231 - 26. In other words, you don’t
need to worry if your program eventually breaks if the user chooses a value for k that’s too big or almost too big to fit in an int. (Recall that an int can overflow.)
But, even if k is greater than 26, alphabetical characters in your program’s input should remain alphabetical characters in your program’s output. For instance, if k is 27,
A should not become [ even though [ is 27 positions away from A in ASCII, per asciichart.com; A should become B, since B is 27 positions away from A, provided you wrap around from Z to A.
Your program must output plaintext: (without a newline) and then prompt the user for a string of plaintext (using get_string).
Your program must output ciphertext: (without a newline) followed by the plaintext’s corresponding ciphertext, with each alphabetical character in the plaintext
"rotated" by k positions; non-alphabetical characters should be outputted unchanged.
Your program must preserve case: capitalized letters, though rotated, must remain capitalized letters; lowercase letters, though rotated, must remain lowercase letters.
After outputting ciphertext, you should print a newline. Your program should then exit by returning 0 from main.
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	offset, err := args()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print("plaintext: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		ciphertext := cipher(input, offset)
		fmt.Printf("ciphertext: %s", ciphertext)
		break
	}
}

//args enforces the number of arguments provided at the CLI and converts user input to an integer
func args() (int, error) {
	var err error
	args := os.Args
	if len(args) <= 1 || len(args) > 2 {
		err := errors.New("Only 1 argument is accepted")
		return 0, err
	}
	//convert string to integer
	offset, err := strconv.Atoi(args[1])
	if err != nil {
		err := errors.New("Only integers are accepted")
		return 0, err
	}
	if offset < 0 {
		err := errors.New("Only non-negative integers are accepted")
		return 0, err
	}
	return offset, nil
}

func cipher(plaintext string, key int) string {
	/*
		Your program must output ciphertext: (without a newline) followed by the plaintext’s corresponding ciphertext, with each alphabetical character in the plaintext
		"rotated" by k positions; non-alphabetical characters should be outputted unchanged.
		Your program must preserve case: capitalized letters, though rotated, must remain capitalized letters; lowercase letters, though rotated, must remain lowercase letters.
		After outputting ciphertext, you should print a newline. Your program should then exit by returning 0 from main.
	*/
	cipherText := ""

	for _, letter := range plaintext {
		lowercase := letter >= 97 && letter <= 122
		uppercase := letter >= 65 && letter <= 90

		var asciia rune
		asciia = 97
		var asciiA rune
		asciiA = 65

		if lowercase {
			offset := (letter - asciia + rune(key)) % 26
			cipherText += string(offset + asciia)
		} else if uppercase {
			offset := (letter - asciiA + rune(key)) % 26
			cipherText += string(offset + asciiA)
		} else {
			//return any character as is.
			cipherText += string(letter)
		}
	}
	return cipherText
}
