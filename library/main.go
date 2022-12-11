package main

import (
	"flag"
	"fmt"
	"library/open"
	"os"
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	getIsbn := getCmd.String("isbn", "", "Book by ISBN")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getIsbn)
	}
}

func HandleGet(getCmd *flag.FlagSet, isbn *string) {
	getCmd.Parse(os.Args[2:])

	if *isbn == "" {
		fmt.Println("requires isbn")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	strPointerValue := *isbn
	if len(strPointerValue) != 13 {
		fmt.Println("isbn must have 13 characters")
		os.Exit(1)
	}

	if len(strPointerValue) == 13 {
		fmt.Println(string(open.GetAuthorFromBookIsbn(isbn)))
	}

}
