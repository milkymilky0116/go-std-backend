package models

import (
	"database/sql"
	"errors"
	"time"
)

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

type GistModel struct {
	DB *sql.DB
}

func (m *GistModel) FindOne(id int) (*Gist, error) {
	var result Gist
	var gistId, writerId int
	var title, content string
	var created_at, updated_at time.Time
	statement := "SELECT * FROM gist WHERE id = ?"
	row := m.DB.QueryRow(statement, id)

	if err := row.Scan(&gistId, &title, &content, &writerId, &created_at, &updated_at); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &result, sql.ErrNoRows
		}
		return &result, err
	}
	var userModel = &UserModel{DB: m.DB}
	user, err := userModel.findUser(writerId)
	if err != nil {
		return &result, err
	}
	result = Gist{
		Id:        gistId,
		Title:     title,
		Content:   content,
		Writer:    *user,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
	return &result, nil
}

func (m *GistModel) FindMany(limit int) ([]*Gist, error) {
	gists := []*Gist{}
	statement := "SELECT id FROM gist ORDER BY created_at DESC LIMIT ?"
	rows, err := m.DB.Query(statement, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gistId int
		gist := &Gist{}
		if err := rows.Scan(&gistId); err != nil {
			return nil, err
		}
		gist, err = m.FindOne(gistId)
		if err != nil {
			return nil, err
		}
		gists = append(gists, gist)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return gists, err
}

func (m *GistModel) Insert(params GistParam) (int, error) {
	statement := "INSERT INTO gist (title, content, writer) VALUES (?, ?, ?)"
	result, err := m.DB.Exec(statement, params.Title, params.Content, params.Writer)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
