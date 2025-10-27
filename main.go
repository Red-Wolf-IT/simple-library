package main

import (
	"fmt"
	"simple-library/library"
)

func main() {
	lib := library.New()

	lib.AddBook("1984", "Джордж Оруэлл", 1949)
	lib.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

	lib.AddReader("Иван", "Иванович", "+79280000000")
	lib.AddReader("Антон", "Сергеевич", "+79288008080")

	b, r, err := lib.IssueBookToReader(1, 1)

	if err != nil {
		fmt.Println("Ошибка выдачи", err)
	} else {
		book, _ := lib.FindBookByID(b)
		reader, _ := lib.FindReaderByID(r)
		fmt.Printf("Книга %s выдана читателю %s.\n", book.GetBookInfo(), reader.GetFullName())
	}

	book, _ := lib.FindBookByID(1)
	if book != nil {
		fmt.Println("Статус книги после выдачи:\n", book)
	}
	
	


	// fmt.Println("Загрузка системы управления библиотекой...")

	// myLibrary := &library.New{}

	// fmt.Println("\n--- Заполнение библиотеки ---")

	// myLibrary.AddReader("Иван", "Иванович", "+79280000000")
	// myLibrary.AddReader("Антон", "Сергеевич", "+79288008080")

	// myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	// myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

	// fmt.Println("\n--- Библиотека готова к работе ---")
	// fmt.Println("Количество читателей:", len(myLibrary.Readers))
	// fmt.Println("Количество книг:", len(myLibrary.Books))

	// fmt.Println("--- Тестируем выдачу книг ---")

	// fmt.Println("Успешная выдача книги:")
	// b, r, err := myLibrary.IssueBookToReader(1, 1)
	// if err != nil {
	// 	fmt.Println("Ошибка выдачи", err)
	// } else {
	// 	book, _ := myLibrary.FindBookByID(b)
	// 	reader, _ := myLibrary.FindReaderByID(r)
	// 	fmt.Printf("Книга %s выдана читателю %s.", book, reader)
	// }

	// book, _ := myLibrary.FindBookByID(1)
	// if book != nil {
	// 	fmt.Println("Статус книги после выдачи:", book)
	// }

	// b, r, err = myLibrary.IssueBookToReader(12, 1)
	// if err != nil {
	// 	fmt.Println("Ожидаемая ошибка:", err)
	// } else {
	// 	book, _ := myLibrary.FindBookByID(b)
	// 	reader, _ := myLibrary.FindReaderByID(r)
	// 	fmt.Printf("Книга %s выдана читателю %s.", book, reader)
	// }

	// book, err = myLibrary.FindBookByID(1)
	// if err != nil {
	// 	fmt.Println("Не удалось найти книгу с таким ID")
	// } else {
	// 	fmt.Println("Статус книги после выдачи:", book)
	// }
}
