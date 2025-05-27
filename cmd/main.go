package main

import (
	"fmt"

	"github.com/lawsonredeye/lms/internal/repository"
	"github.com/lawsonredeye/lms/internal/service"
)

func main() {
	db := repository.NewBookDatabase()
	bookService := service.NewBookService(db)


	id := bookService.CreateBook("1984", "George Orwell", 1949, "Dystopian")
	fmt.Println("Book created with ID:", id)

	bookService.PrintBooks()

}
