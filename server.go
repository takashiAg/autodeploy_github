package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var PORT = 8080

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r)
	//fmt.Fprint(w, 100)

	// header
	method := r.Method
	fmt.Println("[method] " + method)
	for k, v := range r.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// GET
	if method == "POST" {
		r.ParseForm()
		for k, v := range r.Form {
			fmt.Print("[param] " + k)
			fmt.Println(": " + strings.Join(v, ","))
		}
		fmt.Fprint(w, "Recieved Get request!!")
	}
}

func main() {
	port := strconv.Itoa(PORT)
	fmt.Println("access http://127.0.0.1:" + port)
	http.HandleFunc("/", viewHandler) // hello
	http.ListenAndServe(":"+port, nil)
}
