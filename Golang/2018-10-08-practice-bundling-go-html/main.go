package main

import (
	"fmt"
	"net/http"
)

const myindex string = `<!DOCTYPE html>
<html>
    <head>
        <script>
        function myFunction(){
            document.getElementById("demo").innerHTML = "I'm inside of a Go server!";
        }
        function myOtherFunction(){
            document.getElementById("demo").innerHTML = "I'm also inside of a Go server :)";
        }
        </script>
    </head>
    <body>
        <h1>This is being run on a Go server at port 8080!</h1>
        <p id="demo">Below is some embedded javascript functionality inside of the tag</p>
        <button type="button" onclick="myFunction()">Click me to change paragraph with id "demo"</button>
        <button type="button" onclick="myOtherFunction()">Click me to change paragraph with id "demo"</button>
    </body>
</html>`

func serveHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, myindex)
}

func main() {
	http.HandleFunc("/", serveHTML)
	http.ListenAndServe(":8080", nil)
}
