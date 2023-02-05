package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sk25469/zocket-crud-api/pkg/model"
	"github.com/sk25469/zocket-crud-api/pkg/utils"
)

var allBooks []model.Book

func GetAllBook(ctx *fiber.Ctx) error {
	return ctx.JSON(allBooks)
}

func AddNewBook(ctx *fiber.Ctx) error {
	var book model.Book
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	book.Id = uuid.NewString()
	allBooks = append(allBooks, book)
	return ctx.Status(201).JSON(book)
}

func UpdateBookById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	index := utils.GetBookIndex(allBooks, idParam)
	if index == -1 {
		return ctx.Status(501).SendString("No book with this id")
	}
	var updatedBook model.Book
	if err := ctx.BodyParser(&updatedBook); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	if updatedBook.Count != 0 {
		allBooks[index].Count = updatedBook.Count
	}
	if updatedBook.Name != "" {
		allBooks[index].Name = updatedBook.Name
	}
	if updatedBook.Publication != "" {
		allBooks[index].Publication = updatedBook.Publication
	}
	return ctx.Status(200).JSON(allBooks)
}

func DeleteBookById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	index := utils.GetBookIndex(allBooks, idParam)

	if index == -1 {
		return ctx.Status(501).SendString("No book with this id")
	}

	var allBooksAfterDeleting []model.Book

	var err error
	if allBooksAfterDeleting, err = utils.DeleteBook(allBooks, index); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	allBooks = allBooksAfterDeleting
	log.Printf("After deleting: %v", allBooks)

	return ctx.Status(200).JSON("Book deleted successfully")
}
