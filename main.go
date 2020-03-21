package main

import (
	// "bytes"

	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"

	Ascify "./docs/ascify"
	// "strconv"
	// "time"
)

// Data is a struct that will be sent as a respond
type Data struct {
	Output    string
	ErrorNum  int
	ErrorText string
}

var temp *template.Template

// func init() {

// }
func main() {
	http.HandleFunc("/", serverHandler)
	// Creating a handler for handling static files
	FileServer := http.FileServer(http.Dir("docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", FileServer))
	fmt.Println("Server is listening to port #8080 ... ")
	http.ListenAndServe(":8080", nil) // When nil, the DefaultServeMux is used
}

// serverHandler is my very first Handler
func serverHandler(res http.ResponseWriter, req *http.Request) {
	d := Data{}
	temp = template.Must(template.ParseGlob("docs/htmlTemplates/*.html"))

	if req.URL.Path != "/" {
		d.ErrorNum = 404
		d.ErrorText = "Page Not Found"
		errorHandler(res, req, &d) // 404 ERROR
		return
	}

	if req.Method == "GET" {
		temp.ExecuteTemplate(res, "index.html", d)

	} else if req.Method == "POST" {
		// Gathering information to be processed
		text := req.FormValue("input")
		font := req.FormValue("font")

		out, err := Ascify.AsciiArt(text, font)
		if err {
			d.ErrorNum = 500
			d.ErrorText = "Internal Server Error"
			errorHandler(res, req, &d)
			return
		}
		d.Output = out

		if req.FormValue("process") == "show" {
			temp.ExecuteTemplate(res, "index.html", d)
		} else if req.FormValue("process") == "download" {
			a := strings.NewReader(d.Output)
			res.Header().Set("Content-Disposition", "attachment; filename=file.txt")
			res.Header().Set("Content-Length", strconv.Itoa(len(d.Output)))
			io.Copy(res, a)
		} else {
			d.ErrorNum = 400
			d.ErrorText = "Bad Request"
			errorHandler(res, req, &d)
			return
		}
	} else {
		d.ErrorNum = 400
		d.ErrorText = "Bad Request"
		errorHandler(res, req, &d)
		return
	}
}

func errorHandler(res http.ResponseWriter, req *http.Request, d *Data) {
	res.WriteHeader(d.ErrorNum)
	temp.ExecuteTemplate(res, "error.html", d)
}
