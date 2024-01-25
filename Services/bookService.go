package Services

import (
	"BookStoreApi-Go/Model"
	"encoding/json"
	"github.com/google/uuid"
	"sync"
)

var bookList map[uuid.UUID]Model.Book
var IsUuidExist map[uuid.UUID]bool
var mu sync.Mutex

func init() {
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

func SaveBook(book Model.Book) []byte {
	mu.Lock()
	bookList[book.UUID] = book
	mu.Unlock()
	jsonBook, _ := json.Marshal(book)
	return jsonBook
}

func GetBookById(id uuid.UUID) []byte {
	mu.Lock()
	jsonBook, _ := json.Marshal(bookList[id])
	mu.Unlock()
	return jsonBook
}

func GetAllBooks() []byte {
	mu.Lock()
	jsonBook, _ := json.Marshal(bookList)
	mu.Unlock()
	return jsonBook
}
func DeleteBookById(id uuid.UUID) []byte {
	mu.Lock()
	jsonBook, _ := json.Marshal(bookList[id])
	delete(bookList, id)
	IsUuidExist[id] = false
	mu.Unlock()
	return jsonBook
}

func UpdateBookById(id uuid.UUID, book Model.Book) []byte {
	mu.Lock()
	delete(bookList, id)
	book.UUID = id
	bookList[id] = book
	mu.Unlock()
	jsonBook, _ := json.Marshal(book)
	return jsonBook
}
