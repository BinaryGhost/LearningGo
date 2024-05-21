package helpers

import (
	"flag"
	"fmt"
)

// This is an outline of the interface for this URLshortener

func CLI() {
	fmt.Println("Moin Loide")

	create := flag.String("short", "dofunction", "Will create a shortURL and store it in the local DB")
	get := flag.String("get", "dofunction", "Prints to the terminal what URL the shortURL is mapped to")
	replace := flag.String("replace", "dofunction", "Replace a shortURL with another")
	delete := flag.String("delete", "dofunction", "Delete the shortURL from the local DB")

	fmt.Printf("%s, %s, %s, %s", *create, *get, *replace, *delete)
	flag.Parse()
}
