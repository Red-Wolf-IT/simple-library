package domain

import (
	"fmt"
)


type Reader struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	IsActive  bool
}

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int
}



// Геттеры
func (r *Reader) GetFullName() string {
	return r.FirstName + " " + r.LastName
}

func (r *Reader) GetFullInfo() string {
	return r.FirstName + " " + r.LastName + " (" + r.Phone + ")"
}

func (b *Book) GetBookInfo() string {
	return (b.Title + "(" + fmt.Sprintf("%d", b.Year) + ", " + b.Author + ")")
}

// Получение информации о читателе и его статуса
func (r *Reader) DisplayReader() string {
	return fmt.Sprintf("Читатель: %s (ID: %d)(Status: %v)\n", r.GetFullName(), r.ID, r.IsActive)
}

// Переключение статуса
func (r *Reader) Deactivate() {
	r.IsActive = false
}

// Переключение статуса
func (r *Reader) Activate() {
	r.IsActive = true
}

// Получение информации о текущем статусе читателя
func (r *Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен"
	} else {
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь %s, ID: %d, пользователь %s", r.GetFullName(), r.ID, status)
}

// Получение информации о текущем статусе книги
func (b *Book) String() string {
	status := ""
	if b.IsIssued {
		status = "используется"
		return fmt.Sprintf("ID: %d, %s, книга %s читателем %d", b.ID, b.GetBookInfo(), status, *b.ReaderID)
	} else {
		status = "не используется."
		return fmt.Sprintf("ID: %d, %s, книга %s", b.ID, b.GetBookInfo(), status)
	}
}

// Привязка книги к читателю
func (b *Book) IssueBook(r *Reader) (string, error) {
	if b.IsIssued {
		return "", fmt.Errorf("книга %s уже используется", b.GetBookInfo())
	}
	b.IsIssued = true
	b.ReaderID = &r.ID
	return fmt.Sprintf("книга выдана читателю %s.\n", r.GetFullName()), nil
}

// Отвязка книги от читателя
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга %s уже в библиотеке", b.GetBookInfo())
	}

	b.IsIssued = false
	b.ReaderID = nil

	return nil
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s взял книгу %s\n", r.GetFullName(), b.GetBookInfo())
}
