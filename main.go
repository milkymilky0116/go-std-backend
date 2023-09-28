package main

import (
	"database/sql"
	"fmt"

	"github.com/milkymilky0116/go-std-backend/db"
)

func main() {
	newUser := db.CreateAuthor("Milky", sql.NullString{String: "This is test.", Valid: true})
	secondUser := db.CreateAuthor("James", sql.NullString{String: "go with sqlc", Valid: true})
	fmt.Println(db.GetAuthor(newUser.ID).Name)
	fmt.Println(db.GetAuthor(secondUser.ID).Name)
	updatedUser := db.UpdateAuthor(db.UpdateAuthorParams{
		Name: "Kim",
		Bio:  newUser.Bio,
		ID:   newUser.ID,
	})
	fmt.Println(updatedUser.Name)
	users := db.ListAuthors()

	for _, user := range users {
		fmt.Printf("ID: %d - Name: %s - Bio: %s\n", user.ID, user.Name, user.Bio.String)
	}

}
