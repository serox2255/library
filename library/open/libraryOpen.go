package open

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"library/model/entity"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func GetAuthorFromBookIsbn(bookIsbn *string) []byte {
	strPointerValue := *bookIsbn
	str, err := json.Marshal(getBookAuthor(strPointerValue))
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	split := strings.Split(string(str), "\"")

	author, err := http.Get("https://openlibrary.org" + split[13] + ".json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	authorData, err := io.ReadAll(author.Body)
	if err != nil {
		log.Fatal(err)
	}

	authorBooks, err := http.Get("https://openlibrary.org" + split[13] + "/works" + ".json?limit=500")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	authorBooksData, err := io.ReadAll(authorBooks.Body)
	if err != nil {
		log.Fatal(err)
	}

	var obj entity.Author
	json.Unmarshal([]byte(authorData), &obj)
	json.Unmarshal([]byte(authorBooksData), &obj)

	sort.Slice(obj.AuthorWork, func(i, j int) bool {
		return obj.AuthorWork[i].Title < obj.AuthorWork[j].Title
	})
	out, err := yaml.Marshal(&obj)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func getBookAuthor(isbn string) entity.Book {
	book, err := http.Get("https://openlibrary.org/isbn/" + isbn + ".json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	bookData, err := io.ReadAll(book.Body)
	if err != nil {
		log.Fatal(err)
	}

	var obj entity.Book
	json.Unmarshal([]byte(bookData), &obj)

	return obj
}
