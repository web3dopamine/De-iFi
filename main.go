package main

import (
	"log"
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"html/template"
)

//Global variables
	var tpl *template.Template
	
	func init(){
	
		tpl = template.Must(template.New("main").ParseGlob("assets/*.html"))
	}

	func main(){

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)

		

		http.HandleFunc("/", index)

		http.HandleFunc("/css/", serveResource)
    	http.HandleFunc("/vendor/", serveResource)
    	http.HandleFunc("/icons/", serveResource)
    	http.HandleFunc("/js/", serveResource)
    	http.HandleFunc("/fonts/", serveResource)
    	http.HandleFunc("/images/", serveResource)

		http.ListenAndServe(":8090", nil)
		
	}

	func serveResource(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Connected")
	    path := "assets"+r.URL.Path
	    fmt.Println(path)
	    var contentType string
	    if strings.HasSuffix(path, ".css") {
	        contentType = "text/css"
	    } else if strings.HasSuffix(path, ".js") {
	        contentType = "text/javascript"
	    } else if strings.HasSuffix(path, ".png") {
	        contentType = "image/png"
	    } else {
	        contentType = "text/plain"
	    }

	    f, err := os.Open(path)

	    if err == nil {
	        defer f.Close()
	        w.Header().Add("Content-Type", contentType)

	        br := bufio.NewReader(f)
	        br.WriteTo(w)
	    } else {
	        w.WriteHeader(404)
	    }
	}
	func index(w http.ResponseWriter, r *http.Request) {
		
		// w.Header().Set("Content-Type", "text/html")
		tpl.ExecuteTemplate(w, "index.html", nil)
         
	}