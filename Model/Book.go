package Model

import "github.com/google/uuid"

type Book struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	PublishDate string    `json:"publishDate"`
	ISBN        string    `json:"isbn"`
}
