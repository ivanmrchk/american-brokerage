package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Nothing here, but very soon.")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 Page not found")
	}
}

//goapp deploy -application american-brokerage-158918 -version 1 .
