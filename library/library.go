package library

import (
	"fmt"
	"errors"
	"simple-library/domain"
)

type Library struct {
	Books   []*domain.Book
	Readers []*domain.Reader

	lastBookID   int
	lastReaderID int
}

func (lib *Library) GetReaderFullNameByID(id int) (string, error) {
	for _, r := range lib.Readers {
		if r.ID == id {
			return r.FirstName + " " + r.LastName, nil
		}
	}
	return "", fmt.Errorf("человек с ID %d не найден", id)
}

func (lib *Library) GetReaderFullInfoByID(id int) (string, error) {
	for _, r := range lib.Readers {
		if r.ID == id {
			return r.FirstName + " " + r.LastName + " (" + r.Phone + ")", nil
		}
	}
	return "", fmt.Errorf("человек с ID %d не найден", id)
}

func (lib *Library) GetBookInfoByID(id int) (string, error) {
	for _, b := range lib.Books {
		if b.ID == id {
			return b.Title + "(" + fmt.Sprintf("%d", b.Year) + ", " + b.Author + ")", nil
		}
	}
	return "", fmt.Errorf("книга с ID %d не найдена", id)
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

	newReader := &domain.Reader{
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

	newBook := &domain.Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}

	lib.Books = append(lib.Books, newBook)
}

// Нахождение книги по id
func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for i := range lib.Books {
		if lib.Books[i].ID == id {
			return lib.Books[i], nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

// Нахождение читателя по id
func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for i := range lib.Readers {
		if lib.Readers[i].ID == id {
			return lib.Readers[i], nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

// Выдача книги читателю
func (lib *Library) IssueBookToReader(bookID, readerID int) (int, int, error) {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return 0, 0, err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return 0, 0, err
	}

	book.IssueBook(reader)
	return bookID, readerID, nil
}

// Возврат книги в библиотеку
func (lib *Library) ReturnBook(bookID int) error {
	b, err := lib.FindBookByID(bookID)
	if err != nil {
		b.ReturnBook()
	}

	return fmt.Errorf("книга %s не найдена", b.GetBookInfo())
}

// Фабричная функция
func New() *Library {
	l := &Library{}
	return l
}