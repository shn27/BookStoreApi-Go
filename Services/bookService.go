package Services

import (
	"BookStoreApi-Go/Model"
	"encoding/json"
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
	res.Write(jsonBook)
}

func GetBookById(id uuid.UUID, res http.ResponseWriter) {
	mu.Lock()
	jsonBook, _ := json.Marshal(bookList[id])
	mu.Unlock()
	res.Write(jsonBook)
}

func GetAllBooks(res http.ResponseWriter) {
	mu.Lock()
	jsonBook, _ := json.Marshal(bookList)
	mu.Unlock()
	res.Write(jsonBook)
}
func DeleteBookById(id uuid.UUID, res http.ResponseWriter) {
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

	jsonBook, _ := json.Marshal(book)
	res.Write(jsonBook)
}
