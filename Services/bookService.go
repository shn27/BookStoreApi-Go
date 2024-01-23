package Services

import (
	"BookStoreApi-Go/Model"
	"encoding/json"
	"fmt"
	"net/http"
)

var bookList map[string]Model.Book

func Init() {
	bookList = make(map[string]Model.Book)
}

func IsExist(id string) bool {
	if bookList[id].UUID == "" {
		return false
	}
	return true
}

func SaveBook(book Model.Book, res http.ResponseWriter) {
	bookList[book.UUID] = book
	jsonBook, _ := json.Marshal(book)
	fmt.Fprint(res, "Book created successfully\n")
	res.WriteHeader(201)
	res.Write(jsonBook)
}

func GetBookById(id string, res http.ResponseWriter) {
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
func DeleteBookById(id string, res http.ResponseWriter) {
	fmt.Fprint(res, "Book deleted successfully\n")
	jsonBook, _ := json.Marshal(bookList[id])
	res.Write(jsonBook)
	delete(bookList, id)
}

func UpdateBookById(id string, book Model.Book, res http.ResponseWriter) {
	delete(bookList, id)
	bookList[id] = book
	fmt.Fprint(res, "Book updated successfully\n")
	jsonBook, _ := json.Marshal(book)
	res.Write(jsonBook)
}
