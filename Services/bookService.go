package Services

import (
	"BookStoreApi-Go/Model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

var bookList map[uuid.UUID]Model.Book
var IsUuidExist map[uuid.UUID]bool

func Init() {
	bookList = make(map[uuid.UUID]Model.Book)
	IsUuidExist = make(map[uuid.UUID]bool)
}

func IsExist(id uuid.UUID) bool {
	if IsUuidExist[id] == true {
		return true
	}
	return false
}

func SaveBook(book Model.Book, res http.ResponseWriter) {
	bookList[book.UUID] = book
	jsonBook, _ := json.Marshal(book)
	fmt.Fprint(res, "Book created successfully\n")
	res.WriteHeader(201)
	res.Write(jsonBook)
}

func GetBookById(id uuid.UUID, res http.ResponseWriter) {
	jsonBook, _ := json.Marshal(bookList[id])
	res.Write(jsonBook)
	//fmt.Fprint(res, "Name : ", bookList[id].Name, "\nAuthor: ", bookList[id].Author, "\n", "publishDate: ", bookList[id].PublishDate)
}

func GetAllBooks(res http.ResponseWriter) {
	jsonBook, _ := json.Marshal(bookList)
	res.Write(jsonBook)

	//for id, _ := range bookList {
	//	fmt.Fprint(res, "Name : ", bookList[id].Name, "\nAuthor: ", bookList[id].Author, "\n", "publishDate: ", bookList[id].PublishDate, "\n\n")
	//}
}
func DeleteBookById(id uuid.UUID, res http.ResponseWriter) {
	fmt.Fprint(res, "Book deleted successfully\n")
	jsonBook, _ := json.Marshal(bookList[id])
	res.Write(jsonBook)
	delete(bookList, id)
	IsUuidExist[id] = false
}

func UpdateBookById(id uuid.UUID, book Model.Book, res http.ResponseWriter) {
	delete(bookList, id)
	book.UUID = id
	bookList[id] = book
	fmt.Fprint(res, "Book updated successfully\n")
	jsonBook, _ := json.Marshal(book)
	res.Write(jsonBook)
}
