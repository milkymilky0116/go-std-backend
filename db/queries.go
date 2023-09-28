package db

import (
	"database/sql"
	"log"
)

var queries, newCtx = InitDB()

func ListAuthors() []Author {
	authors, err := queries.ListAuthors(newCtx)
	if err != nil {
		log.Fatal(err)
	}
	return authors
}

func CreateAuthor(name string, bio sql.NullString) Author {

	insertedAuthor, err := queries.CreateAuthor(newCtx, CreateAuthorParams{
		Name: name,
		Bio:  bio,
	})

	if err != nil {
		log.Fatal(err)
	}

	return insertedAuthor
}

func GetAuthor(id int64) Author {
	retrievedAuthor, err := queries.GetAuthor(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return retrievedAuthor
}

func UpdateAuthor(name string, bio sql.NullString) Author {
	updatedUser, err := queries.UpdateAuthor(ctx, UpdateAuthorParams{
		Name: name,
		Bio:  bio,
	})
	if err != nil {
		log.Fatal(err)
	}
	return updatedUser
}
