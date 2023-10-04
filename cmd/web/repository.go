package web

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserParam struct {
	Name string `json:"name"`
}
type Gist struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Writer    User      `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GistParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Writer  int    `json:"writer"`
}

func (app *Application) findGist(id int) (Gist, error) {
	var result Gist
	var gistId, writerId int
	var title, content string
	var created_at, updated_at time.Time
	row := app.DB.QueryRow("SELECT * FROM gist WHERE id = ?", id)

	if err := row.Scan(&gistId, &title, &content, &writerId, &created_at, &updated_at); err != nil {
		if err == sql.ErrNoRows {
			return result, sql.ErrNoRows
		}
		return result, err
	}
	user, err := app.findUser(writerId)
	if err != nil {
		return result, err
	}
	result = Gist{
		Id:        gistId,
		Title:     title,
		Content:   content,
		Writer:    user,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
	return result, nil
}

func (app *Application) findUser(id int) (User, error) {
	var user User
	row := app.DB.QueryRow("SELECT * FROM user WHERE id = ?", id)
	if err := row.Scan(&user.Id, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, sql.ErrNoRows
		}
		return user, err
	}
	return user, nil
}

func (app *Application) listGist() ([]Gist, error) {
	var gists []Gist
	rows, err := app.DB.Query("SELECT id FROM gist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gistId int
		var gist Gist
		if err := rows.Scan(&gistId); err != nil {
			return nil, err
		}
		gist, err = app.findGist(gistId)
		if err != nil {
			return nil, err
		}
		gists = append(gists, gist)
	}
	return gists, err
}

func (app *Application) appendGist(params GistParam) (Gist, error) {
	var newGist Gist
	result, err := app.DB.Exec("INSERT INTO gist (title, content, writer) VALUES (?, ?, ?)", params.Title, params.Content, params.Writer)
	if err != nil {
		return newGist, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return newGist, err
	}
	newGist, err = app.findGist(int(id))
	if err != nil {
		return newGist, err
	}
	return newGist, nil
}
