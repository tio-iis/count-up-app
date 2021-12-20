package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

func main() {
	fmt.Println("start")

	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal("can't read the file, err = %v", err)
		return
	}
	htmlStr = string(data)

	http.HandleFunc("/countup", hanlder)
	http.HandleFunc("/get", countHanlder)
	http.HandleFunc("/getindexhtml", getIndexHTML)

	http.ListenAndServe(":8080", nil)
}

var htmlStr string
var count int

func hanlder(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintln(w, count)
}

func countHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, count)
}

func getIndexHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, htmlStr)
}
