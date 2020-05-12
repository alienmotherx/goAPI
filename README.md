# goAPI
A simple HTTP API written in Go with PostgreSQL database.


## How to Use
* Install PostgreSQL on your system. You can skip this step if already installed.
* Clone this repository.
* Modify `app.go` file, line 21 to configure the PostgreSQL database connection.
* Run `$go build` in terminal to build executable file.
* Run `go build` to build an executable file.
* Run/execute the generated program in terminal/command line.
* Navigate to `http://localhost:8000/api/v1/books` on your REST CLIENT [I USED POSTMAN].

* Available functions are GET POST PUT DELETE.
* Use GET to retreive all created books at `http://localhost:8000/api/v1/books` 
* Use POST to create new book at `http://localhost:8000/api/v1/books` With this format:
```{"name":"A Little History of Philosophy","author":"Neil Warburton","published_at:"2011-01-02"```
* published_at is with format YYYY-MM-DD
* Use PUT to get update/edit entry/book at `http://localhost:8000/api/v1/books/{id}`. id must be an index that exists.
* Use DELETE to delete entry/book at `http://localhost:8000/api/v1/books{id}`. Deletes book at given id.


##### TODO
[ ] Add middelware to validate user input when creating new entry
[ ] Add Authentication Feature
[ ] Fix published_at


