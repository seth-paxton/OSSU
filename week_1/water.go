package main

/*
Week 1 - Assignment 1

Background
Suffice it to say that the longer you shower, the more water you use. But just how much? Even if you have a "low-flow" showerhead, odds are your shower spits out 1.5 gallons of water per minute.
A gallon, meanwhile, is 128 ounces, and so that shower spits out 1.5 × 128 = 192 ounces of water per minute. A typical bottle of water (that you might have for a drink, not a shower), meanwhile, might be 16 ounces.
So taking a 1-minute shower is akin to using 192 ÷ 16 = 12 bottles of water. Taking (more realistically, perhaps!) a 10-minute shower, then, is like using 120 bottles of water. Deer Park, that’s a lot of water!
Of course, bottled water itself is wasteful; best to use reusable containers when you can. But it does put into perspective what’s being spent in a shower!

Specification
Write, in a file called water.c in your ~/workspace/pset1/ directory, a program that prompts the user for the length of his or her shower in minutes (as a positive integer)
and then prints the equivalent number of bottles of water (as an integer)

For simplicity, you may assume that the user will input a positive integer, so no need for error-checking (or any loops) this time! And no need to worry about overflow!
*/

import (
	"fmt"
	"os"
)

func main() {
	var err error
	var i int
	fmt.Println("How many minutes do you shower? ")
	_, err = fmt.Scan(&i)
	if err != nil {
		fmt.Printf("Wrong input: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("You use the equivalent of %d water bottles during a %d minute shower!", 192*i/16, i)
}
