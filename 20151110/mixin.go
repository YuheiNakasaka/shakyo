package main

import (
	"fmt"
)

type Accessor interface {
	GetText() string
	SetText(string)
}

type Document struct {
	text string
}

func (d *Document) GetText() string {
	return d.text
}

func (d *Document) SetText(text string) {
	d.text = text
}

type Page struct {
	Document
	Page int
}

func main() {
	// PageはAccessor intefaceを満たすDocumentを継承しているので
	// Accessorに代入できる
	var acsr Accessor = &Page{}

	acsr.SetText("page")
	fmt.Println(acsr.GetText())

}
