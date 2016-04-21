package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"crypto/md5"
	"io"
	"os"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse argument, you have to call this by yourself
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "<h1>Hello Bunchhieng!<h1>") // send data to client side
}

func YoLo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Path[len("/yolo/"):]
		fmt.Println(id)
		if id == "all" {
			fmt.Println("Show all yolo")
		} else {
			id, err := strconv.Atoi(id)
			fmt.Println(id)
			if err != nil {
				fmt.Println(err)
			} else {
				http.Redirect(w, r, "/yolo/", http.StatusFound)
			}
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("/Users/Bunchhieng/Documents/Bunchhieng/gowork/src/github.com/Bunchhieng/Web/SimpleServer/login.gtpl")
		if err != nil {
			fmt.Println(err)
		}
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprint("%d", h.Sum(nil))
		t.Execute(w, token)
	} else {
		r.ParseForm()
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", r.Form["password"])
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, err := template.ParseFiles("/Users/Bunchhieng/Documents/Bunchhieng/gowork/src/github.com/Bunchhieng/Web/SimpleServer/fileupload.gtpl")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/yolo", YoLo)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
