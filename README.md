# BookStoreApi-Go


**Installation**
-----------------------------------------------------------------
Download Go, and set up their path variables.

**Running the server**
-----------------------------------------------------------------
```git clone git@github.com:shn27/BookStoreApi-Go.git``` </br>
```go install``` </br>
```go run main.go -p {port}```</br>
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


**cURL commands**
-----------------------------------------------------------------
**Login and receive a JWT $TOKEN (give name and password as like as .env file)**
```
curl --location --request POST 'http://localhost:3000/login' \
--header 'Authorization: Basic YWRtaW46MTIzNA=='
```

**Add book**
```
curl --location 'http://localhost:3000/create' \
--header 'Content-Type: application/json' \
--header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.H9BCmFXbBsuDHoZF2bYndpev4MOY4mCpY9GjVey0d6U' \
--data '{
    "name" : "THE PRODUCTIBE MUSLIM" ,
    "author" :"Mohammed A. Faris" ,
    "publishDate": "01-01-2011",
    "isbn" : "978-984-8254-54-7"
}'
```

**Show all books**

```
curl --location 'http://localhost:3000/books' \
--header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.H9BCmFXbBsuDHoZF2bYndpev4MOY4mCpY9GjVey0d6U'
```


**Show book with given {id}**

```curl --location 'http://localhost:3000/books/76eb1080-9016-4d8c-9262-a6cb253f7675' \
--header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.H9BCmFXbBsuDHoZF2bYndpev4MOY4mCpY9GjVey0d6U'
```

**Update book with given {id}**
```
curl --location --request PUT 'http://localhost:3000/books/4a53e853-c866-4e45-bc6d-debe7cfaef9e' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46MTIzNA==' \
--data '{
    "name" : "THE PRODUCTIBE MUSLIM" ,
    "author" :"Mohammed A. Faris" ,
    "publishDate": "01-01-2000",
    "isbn" : "978-984-8254-54-7"
}'
```

**Delete book with given {id}**
```
curl --location --request DELETE 'http://localhost:3000/books/61f7297e-63e0-4d82-9b8a-53bde32451d2'
```
**References**
-----------------------------------------------------------------
https://www.alexedwards.net/blog/working-with-cookies-in-go
https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go#using-a-map-to-generate-json
