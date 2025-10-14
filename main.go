package main

func main() {
	user1 = Reader{
		ID: 1,
		FirstName: "Ivan",
		LastName: "Ivanovich",
		IsActive: true,
	}

	user1.Deactivate()
	user1.DisplayReader()
}