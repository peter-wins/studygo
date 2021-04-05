package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server
func f1(w http.ResponseWriter, r *http.Request) {
	// str := `<h1 style="color:green">Hello 世界!</h1>`
	// w.Write([]byte(str))
	b, err := ioutil.ReadFile("index.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/test/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
