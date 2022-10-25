package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := routes()
	http.ListenAndServe(":8080", mux)
}

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(`/time`, timeHandler)
	mux.HandleFunc(`/name/`, helloServer)
	mux.HandleFunc(`/`, helloWorld)
	//mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	return mux
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[6:])
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}
