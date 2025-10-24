package main

import (
	"errors"
	"fmt"
)

type Library struct {
	Books   []*Book
	Readers []*Reader

	lastBookID   int
	lastReaderID int
}

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
	return (b.Title + "(" + string(b.Year) + ", " + b.Author + ")")
}

func (lib *Library) GetReaderFullNameByID(id int) (string, error) {
	for _, r := range lib.Readers {
		if r.ID == id {
			return r.FirstName + " " + r.LastName, nil
		}
	}
	return "", fmt.Errorf("Человек с ID %s не найден", id)
}

func (lib *Library) GetReaderFullInfoByID(id int) (string, error) {
	for _, r := range lib.Readers {
		if r.ID == id {
			return r.FirstName + " " + r.LastName + " (" + r.Phone + ")", nil
		}
	}
	return "", fmt.Errorf("Человек с ID %s не найден", id)
}

func (lib *Library) GetBookInfoByID(id int) (string, error) {
	for _, b := range lib.Books {
		if b.ID == id {
			return b.Title + "(" + string(b.Year) + ", " + b.Author + ")", nil
		}
	}
	return "", fmt.Errorf("Книга с ID %s не найдена", id)
}

// Добавление нового читателя
func (lib *Library) AddReader(firstName, lastName, phoneNumber string) error {
	if firstName == "" || lastName == "" {
		return errors.New("имя и фамилия не могут быть пустыми")
	}

	for _, r := range lib.Readers {
		if phoneNumber == r.Phone {
			return errors.New("читатель с таким номером телефона уже существует")
		}
	}

	lib.lastReaderID++

	newReader := &Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phoneNumber,
		IsActive:  true,
	}

	lib.Readers = append(lib.Readers, newReader)

	return nil
}

// Добавление новой книги
func (lib *Library) AddBook(title, author string, year int) {
	lib.lastBookID++

	newBook := &Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}

	lib.Books = append(lib.Books, newBook)
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
		return "", fmt.Errorf("Книга %s уже используется.", b.GetBookInfo())
	}
	b.IsIssued = true
	b.ReaderID = &r.ID
	return fmt.Sprintf("Книга выдана читателю %s.\n", r.GetFullName()), nil
}

// Отвязка книги от читателя
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("Книга %s уже в библиотеке.", b.GetBookInfo())
	}

	b.IsIssued = false
	b.ReaderID = nil

	return nil
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s взял книгу %s\n", r.GetFullName(), b.GetBookInfo())
}

// Нахождение книги по id
func (lib *Library) FindBookByID(id int) (*Book, error) {
	for i := range lib.Books {
		if lib.Books[i].ID == id {
			return lib.Books[i], nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

// Нахождение читателя по id
func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	for i := range lib.Readers {
		if lib.Readers[i].ID == id {
			return lib.Readers[i], nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

// Выдача книги читателю
func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	book.IssueBook(reader)
	return nil
}

// Возврат книги в библиотеку
func (lib *Library) ReturnBook(bookID int) error {
	b, err := lib.FindBookByID(bookID)
	if err != nil {
		b.ReturnBook()
	}

	return fmt.Errorf("книга %s не найдена.", b.GetBookInfo())
}
