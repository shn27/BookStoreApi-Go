package Model

type Book struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishDate string `json:"publishDate"`
	//ISBN        string	`json:"ISBN"`
}
