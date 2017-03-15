package main

/*
Implement your program in a file called crack.c in a directory called crack.
Your program should accept a single command-line argument: a hashed password.
If your program is executed without any command-line arguments or with more than one command-line argument, your program should print an error (of your choice)
and exit immediately, with main returning 1 (thereby signifying an error).
Otherwise, your program must proceed to crack the given password, ideally as quickly as possible, ultimately printing the password in the clear followed by \n, nothing more, nothing less, with main returning 0.
Assume that each password has been hashed with C’s DES-based (not MD5-based) crypt function.
Assume that each password is no longer than (gasp) four characters
Assume that each password is composed entirely of alphabetical characters (uppercase and/or lowercase).
*/

import (
	"crypto/des"
	"errors"
	"fmt"
	"os"
)

func main() {
	args, err := arguments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(args)
}

func cracker() {
	key := "12345678" // 8 bytes! this is the DES block size in bytes

	block, err := des.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(block)

}

//args enforces the number of arguments provided at the CLI and converts user input to an integer
func arguments() (string, error) {
	args := os.Args
	if len(args) <= 1 || len(args) > 2 {
		err := errors.New("Only 1 argument is accepted")
		return "", err
	}
	return args[1], nil
}
