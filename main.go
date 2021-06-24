package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handle)

	port := ":8080"
	println("Server listen on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	var err error
	rr := &request{}
	rr.Method = r.Method
	rr.Headers = r.Header
	rr.URL = r.URL.String()
	rr.Body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rrb, err := json.Marshal(rr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(r.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rrb)
}

type request struct {
	URL     string      `json:"url"`
	Method  string      `json:"method"`
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}

// curl -i -X OPTIONS localhost:8080     команда для отправки запроса из консоли

// флаги для запуска программы
// func getParams() string {
// 	fieldPtr := flag.String("field", "field", "a string")
// 	oldPtr := flag.String("old", "old", "a string")
// 	newPtr := flag.String("new", "new", "a string")

// 	flag.Parse()
// }
