package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "strconv"
	"regexp"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("username is empty")
		}
		// getint, err := strconv.Atoi(r.Form.Get("age"))
		// if err != nil {
		// 	fmt.Println("age is empty")
		// }
		// if getint > 100 {
		// 	fmt.Println("age is too large")
		// }
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			fmt.Println("age must be integer")
		}
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
			fmt.Println("username must be English")
		}
		slice := []string{"football", "basketball", "tennis"}
		conveted_slice := make([]interface{}, len(r.Form["interest"]))
		for i, v := range r.Form["interest"] {
			conveted_slice[i] = v
		}
		conveted_slice2 := make([]interface{}, len(slice))
		for i, v := range slice {
			conveted_slice2[i] = v
		}

		a := Slice_diff(conveted_slice, conveted_slice2)
		if a == nil {
			fmt.Println("valid interest")
		} else {
			if str, ok := a[0].(string); ok { // []inteface{}をstringに変換
				fmt.Println("invalid interest " + str)
			}
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("age:", r.Form["age"])
	}
}

func Slice_diff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !In_slice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return diffslice
}

func In_slice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
