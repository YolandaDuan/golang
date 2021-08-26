package main

import (
	"fmt"
	"io/ioutil"
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
	os.Remove("rmPage2.txt") //remove a single file
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "rmPage2", Body: []byte("2222")}
	p1.save()
	p2, _ := loadPage("rmPage2")
	fmt.Println(string(p2.Body))
	remove()
}
