package main

import (
	"fmt"
	"log"

	"github.com/lawsonredeye/lms/internal/repository"
	"github.com/lawsonredeye/lms/internal/service"
)

func main() {
	db := repository.NewBookDatabase()
	bookService := service.NewBookService(db)

	id := bookService.CreateBook("1984", "George Orwell", 1949, "Dystopian")
	fmt.Println("Book created with ID:", id)

	if err := bookService.UpdateBookByID(id, "The great wall of china", "", 0, ""); err != nil {
		log.Fatal(err)
	}

	d, _ := bookService.GetBookByID(id)
	log.Println(d)
	bookService.PrintBooks()
}
