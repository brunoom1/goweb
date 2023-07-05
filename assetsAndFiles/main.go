package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	ID     uint
	Title  string
	Status string
}

type TodoPage struct {
	Title    string
	TodoList []Todo
}

func main() {

	todoPage := TodoPage{
		Title: "Todo Page",
		TodoList: []Todo{
			{ID: 1, Title: "Fazer isso", Status: "para-fazer"},
			{ID: 2, Title: "fazendo isso", Status: "fazendo"},
			{ID: 3, Title: "Finalizado", Status: "feito"},
		},
	}

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/home.html"))

		temp.Execute(w, todoPage)
	})

	http.ListenAndServe(":8000", nil)
}
