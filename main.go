package main

// https://mw2.google.com/mw-panoramio/photos/medium/60435584.jpg

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func director(r *http.Request) {
}

const CACHE_PATH = "./sdcard"

const CDN = "mw2.google.com"

func main() {

	proxy := httputil.ReverseProxy{
		Transport: http.NewFileTransport(http.Dir(CACHE_PATH)),
		Director:  director,
	}

	http.HandleFunc("/", proxy.ServeHTTP)
	log.Println("Running...")
	panic(http.ListenAndServe(":8080", nil))
}

func logPath(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)

		return
		_, err := os.Stat(CACHE_PATH + r.URL.Path)
		if os.IsNotExist(err) {
			log.Println("Miss")
		} else {
			p.ServeHTTP(w, r)
		}

	}
}
