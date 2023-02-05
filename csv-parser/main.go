package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Book struct {
	Id          string
	Name        string
	Publication string
}

func createBookList(data [][]string) []Book {
	var book []Book
	for i, line := range data {
		if i > 0 {
			var rec Book
			for j, field := range line {
				if j == 0 {
					rec.Id = field
				} else if j == 1 {
					rec.Name = field
				} else if j == 2 {
					rec.Publication = field
				}
			}
			book = append(book, rec)
		}
	}
	return book
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Id\tName\tPublication\n")
	bookList := createBookList(data)

	for i := 0; i < len(bookList); i++ {
		fmt.Printf("%v\t%v\t%v\t\n", bookList[i].Id, bookList[i].Name, bookList[i].Publication)
	}
}
