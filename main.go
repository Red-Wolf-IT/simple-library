package main

import "fmt"

func main() {
	// user1 := Reader{
	// 	ID: 1,
	// 	FirstName: "Ivan",
	// 	LastName: "Ivanovich",
	// 	IsActive: true,
	// }

	// book1 := Book{
	// 	ID:       1,
	// 	Year:     1867,
	// 	Title:    "Война и мир",
	// 	Author:   "Лев Толстой",
	// 	IsIssue: false,
	// }

	// fmt.Println(user1)
	// fmt.Println(book1)
	// book1.IssueBook(&user1)
	// fmt.Println(book1)
	// book1.ReturnBook()
	// fmt.Println(book1)
	// user1.AssignBook(&book1)

	// email := EmailNotifier{Email: "test@mail.ru"}
	// sms := SMSNotifier{PhoneNumber: "89289289898"}

	// email.Notify("test")
	// sms.Notify("test")


	fmt.Println("Загрузка системы управления библиотекой...")

	myLibrary := &Library{}

	fmt.Println("\n--- Заполнение библиотеки ---")

	myLibrary.AddReader("Иван", "Иванович")
	myLibrary.AddReader("Антон", "Сергеевич")

	myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

	fmt.Println("\n--- Библиотека готова к работе ---")
	fmt.Println("Количество читателей:", len(myLibrary.Readers))
	fmt.Println("Количество книг:", len(myLibrary.Books))

	fmt.Println("--- Тестируем выдачу книг ---")

	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println("Ошибка выдачи", err)
	}

	book, _ := myLibrary.FindBookByID(1)
	if book != nil {
		fmt.Println("Статус книги после выдачи:", book)
	}

	err = myLibrary.IssueBookToReader(12, 1)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}
}