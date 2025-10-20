package main

import "fmt"

type Library struct {
	Books	[]*Book
	Readers	[]*Reader

	lastBookID		int
	lastReaderID	int
}

func (lib *Library) AddReader(firstName, lastName string) *Reader {
	lib.lastReaderID++

	newReader := &Reader{
		ID:			lib.lastReaderID,
		FirstName:	firstName,
		LastName:	lastName,
		IsActive:	true,
	}

	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Добавлен новый читатель: %s\n", newReader)
	return newReader
}

func (lib *Library) AddBook(title, author string, year int) *Book {
	lib.lastBookID++

	newBook := &Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued:  false,
		ReaderID: new(int),
	}

	lib.Books = append(lib.Books, newBook)
	fmt.Printf("Добавлена новая книга: %s\n", newBook)
	return newBook
}


type Reader struct {
	ID			int
	FirstName	string
	LastName	string
	IsActive	bool
}

func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)(Status: %v)\n", r.FirstName, r.LastName, r.ID, r.IsActive)
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен"
	} else {
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

type Book struct {
	ID			int
	Title 		string
	Author 		string
	Year		int
	IsIssued 	bool
	ReaderID 	*int
}

func (b Book) String() string{
	status := ""
	if b.IsIssued {
		status = "используется"
		return fmt.Sprintf("ID: %d, %s (%s, %d), книга %s читателем %d", b.ID, b.Title, b.Author, b.Year, status, *b.ReaderID)
	} else {
		status = "не используется."
		return fmt.Sprintf("ID: %d, %s (%s, %d), книга %s", b.ID, b.Title, b.Author, b.Year, status)
	}
	
}

func (b *Book) IssueBook(r *Reader) {
	if b.IsIssued{
		fmt.Println("Книга уже используется.")
	} else {
		b.IsIssued = true
		b.ReaderID = &r.ID
		fmt.Printf("Книга выдана читателю %s %s.\n", r.FirstName, r.LastName)
	}
}

func (b *Book) ReturnBook() {
	if !b.IsIssued{
		fmt.Println("Книга уже в библиотеке.")
	} else {
		b.IsIssued = false
		b.ReaderID = nil
		fmt.Println("Книга возвращена.")
	}
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s(%s %d)\n", r.FirstName, r.LastName, b.Title, b.Author, b.Year)
}


func (lib *Library) FindBookByID(id int) (*Book, error) {
	for i := range lib.Books {
		if lib.Books[i].ID == id {
			return lib.Books[i], nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	for i := range lib.Readers {
		if lib.Readers[i].ID == id {
			return lib.Readers[i], nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

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
