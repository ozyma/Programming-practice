package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//Page struct will be used to store a wiki page for us.
type Page struct {
	Title string
	//We are not using a string because the ioutil expect byte slice types.
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	//function returns an error because that is the return type of the ioutil.WriteFile function in package ioutil
	return ioutil.WriteFile(filename, p.Body, 0600)
	/*If all goes well, Page.save() will return nil (the zero-value for pointers, interfaces, and some other types).
	The octal integer literal 0600, passed as the third parameter to WriteFile, indicates that the file should be
	created with read-write permissions for the current user only. (See the Unix man page open(2) for details.)
	*/
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewhandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

func edithandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a simple page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewhandler)
	http.HandleFunc("/edit/", edithandler)
	//http.HandleFunc("/save/", savehandler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
