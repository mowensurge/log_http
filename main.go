package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "port server listens on")
	flag.Parse()
	fmt.Printf("listening on port %d\n", *port)
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(string(b))
}
