package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./http/serve/xx.txt")
	fmt.Println(r)
	if err != nil {
		_, err := w.Write([]byte(fmt.Sprintf("%v", err)))
		if err != nil {
			return
		}
		return
	}
	_, err = w.Write(file)
	if err != nil {
		return
	}
}

func f2(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(r.URL)
	fmt.Println(queryParam)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	_, err := w.Write([]byte("ok!"))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/posts/go", f1)
	http.HandleFunc("/posts/xx", f2)
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		return
	}
}
