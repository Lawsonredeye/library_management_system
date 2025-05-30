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
	// bookService.PrintBooks()

	mDB := repository.NewMemberDatabase() // This initialised the database that would be used to store the members credentials
	membService := service.NewMemberService(mDB)

	loanDB := repository.NewLoanDB()
	loanService := service.NewLoanService(loanDB)

	// Since this is  a CLI application then i would be handling everything here, Sorry ^_^
	membService.CreateMember("lawson", "you cant copy this flow, haha")
	membService.CreateMember("Benson", "you cant copy this flow, haha")
	membService.CreateMember("Denson", "you cant copy this flow, haha")
	res := membService.CreateMember("Dickson", "you cant copy this flow, haha")
	printOut(*membService)

	membService.UpdateMemberByID(res, "", "how on earth is this possible!")
	printOut(*membService)

	if err := bookService.BorrowBook(id, res); err != nil {
		log.Fatalln(err)
	}

	loanService.BorrowAnyBook(id, res) // Can only handle one book at this point, might change that if i use a real db
	loanService.PrintBorrowedBooks()
}

func printOut(membService service.MemberService) {
	resp := membService.GetAllMembers()
	for idx, val := range resp {
		fmt.Printf("member %v: %+v\n", idx, *val)
	}
}
