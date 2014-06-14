package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

const MYFANCYPASSWORD = "AAABBB*@(&^&%^&!*@_O_}PLDKHOUGYQ&T@(*HIPNCK:nobueifvwyqg9h029pn3"
const MYSQL_USER = "cpt_ahab"

func director(r *http.Request) {
}

const path = http.Dir("./sdcard")

func main() {
	fmt.Println(MYSQL_USER)

	proxy := httputil.ReverseProxy{
		Transport: http.NewFileTransport(path),
		Director:  director,
	}

	http.HandleFunc("/", handler(&proxy))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.Header().Set("X-Ben", MYFANCYPASSWORD)
		p.ServeHTTP(w, r)
	}
}
