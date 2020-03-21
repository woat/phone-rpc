package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	http.HandleFunc("/rpc/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("rpc hit")
		ioutil.WriteFile("/Users/addy/Desktop/doodoo", []byte("poopoo"), 0644)
	})
	http.ListenAndServe(":8080", nil)
}
