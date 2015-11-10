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

type Article struct {
}

func SetAndGet(acsr Accessor) {
	acsr.SetText("accessor")
	fmt.Println(acsr.GetText())
}

func main() {
	SetAndGet(&Page{})
	SetAndGet(&Document{})
	// SetAndGet(&Article{}) => ArticleはAccessorのinterfaceを実装していないのでエラー
}
