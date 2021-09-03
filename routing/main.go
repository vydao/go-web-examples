package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page, err := strconv.Atoi(vars["page"])

		if err == nil {
			fmt.Fprintf(w, "You have requested the book: %s on page %d\n", title, page)
		}
	})

	http.ListenAndServe(":8080", r)
}
