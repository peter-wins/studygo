package main

import "net/http"

// net/http server
func f1(w http.ResponseWriter, r *http.Request) {
	str := "<h4> Hello world! </h4>"
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/test", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
