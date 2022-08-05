package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w httprouter.ResponseWriter, r *httprouter.Request, _ httprouter.Params) {
		log.Println("GET /")
	})

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
