package Services

import (
	"BookStoreApi-Go/Model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

var bookList map[uuid.UUID]Model.Book
var IsUuidExist map[uuid.UUID]bool
var mu sync.Mutex

func Init() {
	bookList = make(map[uuid.UUID]Model.Book)
	IsUuidExist = make(map[uuid.UUID]bool)
}

func IsExist(id uuid.UUID) bool {
	mu.Lock()
	if IsUuidExist[id] == true {
		mu.Unlock()
		return true
	}
	mu.Unlock()
	return false
}

func SaveBook(book Model.Book, res http.ResponseWriter) {
	mu.Lock()
	bookList[book.UUID] = book
	mu.Unlock()
	jsonBook, _ := json.Marshal(book)
	fmt.Fprint(res, "Book created successfully\n")
	res.WriteHeader(201)
	res.Write(jsonBook)
}

func GetBookById(id uuid.UUID, res http.ResponseWriter) {
	mu.Lock()
	jsonBook, _ := json.Marshal(bookList[id])
	mu.Unlock()
	res.Write(jsonBook)
	//fmt.Fprint(res, "Name : ", bookList[id].Name, "\nAuthor: ", bookList[id].Author, "\n", "publishDate: ", bookList[id].PublishDate)
}

func GetAllBooks(res http.ResponseWriter) {
	mu.Lock()
	jsonBook, _ := json.Marshal(bookList)
	mu.Unlock()
	res.Write(jsonBook)

	//for id, _ := range bookList {
	//	fmt.Fprint(res, "Name : ", bookList[id].Name, "\nAuthor: ", bookList[id].Author, "\n", "publishDate: ", bookList[id].PublishDate, "\n\n")
	//}
}
func DeleteBookById(id uuid.UUID, res http.ResponseWriter) {
	fmt.Fprint(res, "Book deleted successfully\n")

	mu.Lock()
	jsonBook, _ := json.Marshal(bookList[id])
	res.Write(jsonBook)
	delete(bookList, id)
	IsUuidExist[id] = false
	mu.Unlock()
}

func UpdateBookById(id uuid.UUID, book Model.Book, res http.ResponseWriter) {
	mu.Lock()
	delete(bookList, id)
	book.UUID = id
	bookList[id] = book
	mu.Unlock()

	fmt.Fprint(res, "Book updated successfully\n")
	jsonBook, _ := json.Marshal(book)
	res.Write(jsonBook)
}
