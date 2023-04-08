package main

import (
	"github.com/atorahmad/go/crud/initializers"
	"github.com/atorahmad/go/crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.KonekDB()
}

func main() {
	r := gin.Default()

	//books routes
	routes.BookRoutes(r)

	r.Run()
}

//postman
//post a book: localhost:3001/books
//get books: localhost:3001/books?page=1&limit=10&search=To Kill a Mockingbird&sort=year&year=1960
//get by id: localhost:3001/books/15
//put a book: localhost:3001/books/1
//delete a book: localhost:3001/books/19
