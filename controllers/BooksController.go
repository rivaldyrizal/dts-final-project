package BooksController

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	Model "project/models"
)

func sqliteDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Fprint(w, "runing")

	db := sqliteDb()

	var Books []Model.Books
	db.Find(&Books)

	datas := map[string]interface{}{
		"Books": Books,
	}

	files := []string{
		"views/index.html",
		"views/base.html",
	}

	template, err := template.ParseFiles(files...)

	if err != nil {
		fmt.Println(err)
	}

	template.ExecuteTemplate(w, "base", datas)
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := sqliteDb()

	if r.Method == "POST" {
		Books := Model.Books{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
		}

		db.Create(&Books)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		files := []string{
			"views/create.html",
			"views/base.html",
		}

		template, err := template.ParseFiles(files...)

		if err != nil {
			fmt.Println(err)
		}

		template.ExecuteTemplate(w, "base", nil)
	}
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := sqliteDb()

	id := ps.ByName("id")

	var Books Model.Books
	db.First(&Books, id)

	db.Delete(&Books)
	http.Redirect(w, r, "/", http.StatusFound)
}
