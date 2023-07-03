package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func main () {

	router := mux.NewRouter()
	router.HandleFunc("/book/{title}/page/{page}", func (w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)

		fmt.Fprintf(w, "book %s - page %s", vars["title"], vars["page"])
	})

	http.ListenAndServe(":8080", router)
}