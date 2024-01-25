# BookStoreApi-Go


**Installation**
-----------------------------------------------------------------
Download Go, and set up their path variables.

**Running the server**
-----------------------------------------------------------------
```git clone git@github.com:shn27/BookStoreApi-Go.git``` </br>
```go install```
```go run main.go -p {port}```
Use postman for testing.



**API Endpoints**
-----------------------------------------------------------------
|method|url|body|action
|-----|----|---|---|
|GET| `http://localhost:port/login` | --header 'Authorization: Basic c2FtaToxMjM0' | returns a JWT token $TOKEN into cookies|
|GET| `http://localhost:port/books` |   | returns all the books.|
|GET| `http://localhost:port/books/{id}` |   | return a single book where Id = bookId.|
|POST| `http://localhost:port/create` |   | Add the book. Return the addded book.|
|PUT| `http://localhost:port/books/{id}` |   | Update the book if bookId is present. Return the updated book.|
|DELETE| `http://localhost:port/books/{id}` |   | Delete the book if bookId is present. Return the deleted book.
