package main

import (
	"book_inventory/app"
	"book_inventory/auth"
	"book_inventory/db"
	"book_inventory/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	handler := app.New(db)

	r.GET("/", auth.HomeHandler)
	r.GET("/login", auth.LoginGetHandler)

	r.GET("/books", middleware.AuthValidator, handler.GetBooks)
	r.GET("/book/:id", middleware.AuthValidator, handler.GetBookById)

	r.GET("/addBook", middleware.AuthValidator, handler.GetAddBook)
	r.POST("/book", middleware.AuthValidator, handler.PostBook)

	r.GET("/updateBook/:id", middleware.AuthValidator, handler.GetUpdateBook)
	r.POST("/updateBook/:id", middleware.AuthValidator, handler.PutBook)

	r.POST("/deleteBook/:id", middleware.AuthValidator, handler.DeleteBook)

	r.POST("/login", auth.LoginPostHandler)

	r.Run(":8081")
}
