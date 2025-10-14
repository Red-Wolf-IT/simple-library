package main

import "fmt"

func main() {
	user1 := Reader{
		ID: 1,
		FirstName: "Ivan",
		LastName: "Ivanovich",
		IsActive: true,
	}

	book1 := Book{
		ID:       1,
		Year:     1867,
		Title:    "Война и мир",
		Author:   "Лев Толстой",
		IsIssue: false,
	}

	fmt.Println(user1)
	fmt.Println(book1)
	book1.IssueBook(&user1)
	fmt.Println(book1)
	book1.ReturnBook()
	fmt.Println(book1)
	user1.AssignBook(&book1)
}