package main

import (
	"log"
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"html/template"
	"strconv"
	"math"
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
		http.HandleFunc("/zakaat", zakaat)

		http.HandleFunc("/css/", serveResource)
    	http.HandleFunc("/vendor/", serveResource)
    	http.HandleFunc("/icons/", serveResource)
    	http.HandleFunc("/js/", serveResource)
    	http.HandleFunc("/fonts/", serveResource)
    	http.HandleFunc("/images/", serveResource)

		http.ListenAndServe(":8090", nil)
		
	}

	func serveResource(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Connected")
	    path := "assets"+r.URL.Path
	    // fmt.Println(path)
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

	func zakaat(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		return
		} 

		var (
			amt string
			asset string
		)

		amt = r.FormValue("amt")
		asset = r.FormValue("asset")
		
		fmt.Println("amount = "+amt)
		fmt.Println("asset = "+asset)

		fmt.Println(price(asset))

		balance, err := strconv.ParseFloat(amt, 64)
		asset_price := price(asset)

		if err != nil {
			log.Fatal(err)
		}

		final := balance * asset_price

		

		final = math.Round(final*10)/10
		fmt.Println(final)
		fmt.Fprint(w, final)



		// Malaysian Ringgit
		rm := final * 4.30

		nisaab := 75 * 240.15

		OnZakaat := rm - nisaab

		zakaatinRinggit := OnZakaat * 0.025

		// fmt.Fprint(w, zakaat)
		fmt.Println("Ringgit")
		fmt.Println(rm)
		fmt.Println("nisaab")
		fmt.Println(nisaab)
		fmt.Println("on zakaat")
		fmt.Println(OnZakaat)
		fmt.Println("zakaat in Ringgit")
		fmt.Println(zakaatinRinggit)
		// Malaysian Ringgit
		// zakaat := OnZakaat * 0.25
	}