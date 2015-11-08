package main

import (
	"fmt"
	"net/http"
)

type String string

// 既存の型にSerevHTTPメソッドを生やす
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/", String("hello, world"))
	http.ListenAndServe("localhost:8080", nil)
}
