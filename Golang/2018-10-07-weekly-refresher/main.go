package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, index)
}

func differentIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, newIndex)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", serveIndex)
	go http.ListenAndServe(":8080", r) //our x mux will not run unless this is sent off into its own thread

	x := mux.NewRouter()
	x.HandleFunc("/", differentIndex)
	http.ListenAndServe(":8081", x)
}

//string with an html as a multiline string
const index string = `<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
        
    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 50px;
        background-color: #fff;
        border-radius: 1em;
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        body {
            background-color: #fff;
        }
        div {
            width: auto;
            margin: 0 auto;
            border-radius: 0;
            padding: 1em;
        }
    }
    </style>    
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is established to be used for illustrative examples in documents. You may use this
    domain in examples without prior coordination or asking for permission.</p>
    <p><a href="http://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>`

const newIndex string = `<!DOCTYPE html>
<html>
    <head>
        <script>
        function myFunction(){
            document.getElementById("demo").innerHTML = "Go fuck yourself";
        }
        function myOtherFunction(){
            document.getElementById("demo").innerHTML = "Paragraph changed on button click using myOtherFunction()";
        }
        </script>
    </head>
    <body>
        <h1>A webpage with javascript functionality</h1>
        <p id="demo">A paragraph</p>
        <button type="button" onclick="myFunction()">Click me to change paragraph with id "demo"</button>
        <button type="button" onclick="myOtherFunction()">Click me to change paragraph with id "demo"</button>
    </body>
</html>`
