package main

import "fmt"

type Reader struct {
	ID			int
	FirstName	string
	LastName	string
	IsActive	bool
}

func (r Rider) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
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
	ID		int
	Title 	string
	Author 	string
	IsIssue bool
}