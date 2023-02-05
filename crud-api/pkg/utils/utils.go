package utils

import (
	"github.com/sk25469/zocket-crud-api/pkg/model"
)

func GetBookIndex(allBooks []model.Book, id string) int {
	for i := 0; i < len(allBooks); i++ {
		if allBooks[i].Id == id {
			return i
		}
	}
	return -1
}

func DeleteBook(allBooks []model.Book, index int) ([]model.Book, error) {
	allBooks[index] = allBooks[len(allBooks)-1]
	return allBooks[:len(allBooks)-1], nil

}
