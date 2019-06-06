package main

import (
	"net/http"
	"strings"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/ping", ping)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438
