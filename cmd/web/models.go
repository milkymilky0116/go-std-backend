package web

import "github.com/milkymilky0116/go-std-backend/internal/models"

type TemplateData struct {
	CurrentYear int
	Gist        *models.Gist
	Gists       []*models.Gist
}
