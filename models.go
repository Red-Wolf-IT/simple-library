package main

import "fmt"

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
	IsIssue 	bool
	ReaderID 	*int
}

func (b Book) String() string{
	status := ""
	if b.IsIssue {
		status = "используется."
		return fmt.Sprintf("ID: %d, %s (%s, %d), книга %s читателем %d", b.ID, b.Title, b.Author, b.Year, status, *b.ReaderID)
	} else {
		status = "не используется."
		return fmt.Sprintf("ID: %d, %s (%s, %d), книга %s", b.ID, b.Title, b.Author, b.Year, status)
	}
	
}

func (b *Book) IssueBook(r *Reader) {
	if b.IsIssue{
		fmt.Println("Книга уже используется.")
	} else {
		b.IsIssue = true
		b.ReaderID = &r.ID
		fmt.Printf("Книга выдана читателю %s %s.\n", r.FirstName, r.LastName)
	}
}

func (b *Book) ReturnBook() {
	if !b.IsIssue{
		fmt.Println("Книга уже в библиотеке.")
	} else {
		b.IsIssue = false
		b.ReaderID = nil
		fmt.Println("Книга возвращена.")
	}
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s(%s %d)", r.FirstName, r.LastName, b.Title, b.Author, b.Year)
}