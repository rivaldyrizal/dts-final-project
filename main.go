package main

import (
	"fmt"
	"net/http"

	BooksController "project/controllers"

	Model "project/models"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Model.Books{})
	router := httprouter.New()

	// router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 	fmt.Fprint(w, "Hello World")
	// })

	router.GET("/", BooksController.Index)
	router.GET("/create", BooksController.Create)
	router.POST("/create", BooksController.Create)
	// router.GET("/edit/:id", BooksController.Edit)
	// router.POST("/edit/:id", BooksController.Edit)
	router.GET("/delete/:id", BooksController.Delete)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
