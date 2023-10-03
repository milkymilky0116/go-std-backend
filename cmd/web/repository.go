package web

import "time"

type User struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Gist struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Writer    User      `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var user1 = User{
	Name:      "Milky",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func GistsList() []Gist {
	return []Gist{
		{
			Id:        1,
			Title:     "This is one Gist",
			Content:   "Hi",
			Writer:    user1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        2,
			Title:     "This is two Gist",
			Content:   "Hello",
			Writer:    user1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        3,
			Title:     "This is three Gist",
			Content:   "Hello",
			Writer:    user1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
