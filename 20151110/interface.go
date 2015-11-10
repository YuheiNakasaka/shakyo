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

func main() {
	// SetTextは自身(Document)をポインタで受けプロパティを変更してるので
	// インスタンスも参照渡しで受ける必要あり
	var doc *Document = &Document{}
	doc.SetText("document")
	fmt.Println(doc.GetText())

	// interface型のメソッドを実装しているからinteface型に代入可能
	var acsr Accessor = &Document{}
	acsr.SetText("accessor")
	fmt.Println(acsr.GetText())
}
