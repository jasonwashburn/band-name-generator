package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("What was the name of the stree you grew up on? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	streetName := scanner.Text()
	fmt.Println("What is the name of your pet? ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	petName := scanner.Text()
	fmt.Println("Your band name could be", streetName, petName)
}
