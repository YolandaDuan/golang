package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func remove() {
	os.Remove("rmPage2.txt") //remove a single file rmPage2.txt
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[3:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	p1 := &Page{Title: "rmPage2", Body: []byte("2222")}
	p1.save()
	p2, _ := loadPage("rmPage2")
	fmt.Println(string(p2.Body))
	remove()

	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
