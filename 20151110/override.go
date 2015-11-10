package main

import (
	"fmt"
	"strconv"
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

type ExtendedPage struct {
	Document
	Page int
}

func (ep *ExtendedPage) GetText() string {
	return strconv.Itoa(ep.Page) + " : " + ep.Document.GetText()
}

func main() {
	var acsr Accessor = &ExtendedPage{
		Document{},
		2,
	}
	acsr.SetText("page")
	fmt.Println(acsr.GetText())
}
