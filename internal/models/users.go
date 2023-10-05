package models

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

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) findUser(id int) (*User, error) {
	user := &User{}
	row := m.DB.QueryRow("SELECT * FROM user WHERE id = ?", id)
	if err := row.Scan(&user.Id, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, sql.ErrNoRows
		}
		return user, err
	}
	return user, nil
}
