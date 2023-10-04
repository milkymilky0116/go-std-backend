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

var data = []Gist{
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

var user1 = User{
	Name:      "Milky",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func findGist(id int) Gist {
	var result Gist
	for _, gist := range data {
		if gist.Id == id {
			return gist
		}
	}
	return result
}

func listGist() []Gist {
	return data
}

// func appendGist(title string, content string) {
// 	newGist := Gist{
// 		Title:     title,
// 		Content:   content,
// 		Writer:    user1,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	data = append(data, newGist)
// }
