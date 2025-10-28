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
}
