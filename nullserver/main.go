package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const wait = 1 * time.Second

func OurLoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r.URL)
		log.Println(*r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	fileHandler := http.FileServer(http.Dir("/tmp"))
	wrappedHandler := OurLoggingHandler(fileHandler)

	t := time.NewTimer(wait)
	<-t.C

	tt := time.NewTicker(100 * time.Millisecond)
	for i := 0; i < 2; i++ {
		<-tt.C
	}

	http.ListenAndServe(":8080", wrappedHandler)
}
