package main

import "fmt"

func main() {
	fmt.Println("Загрузка системы управления библиотекой...")

	myLibrary := &Library{}

	fmt.Println("\n--- Заполнение библиотеки ---")

	myLibrary.AddReader("Иван", "Иванович", "+79280000000")
	myLibrary.AddReader("Антон", "Сергеевич", "+79288008080")

	myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

	fmt.Println("\n--- Библиотека готова к работе ---")
	fmt.Println("Количество читателей:", len(myLibrary.Readers))
	fmt.Println("Количество книг:", len(myLibrary.Books))

	fmt.Println("--- Тестируем выдачу книг ---")

	fmt.Println("Успешная выдача книги:")
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println("Ошибка выдачи", err)
	} else {
		fmt.Printf("Книга %s выдана читателю %s.")
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
