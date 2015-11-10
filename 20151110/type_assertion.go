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

type Getter interface {
	GetText() string
}

func dynamicIf(v interface{}) string {
	var result string
	g, ok := v.(Getter)
	if ok {
		result = g.GetText()
	} else {
		result = "not implemented"
	}
	return result
}

func dynamicSwitch(v interface{}) string {
	var result string
	// .(type)はswitchと使わないといけない
	switch checked := v.(type) {
	case Getter:
		result = checked.GetText()
	case string:
		result = "not implemented"
	}
	return result
}

func main() {
	var ep *ExtendedPage = &ExtendedPage{
		Document{},
		3,
	}

	ep.SetText("page")

	fmt.Println(dynamicIf(ep))
	fmt.Println(dynamicIf("string"))

	fmt.Println(dynamicSwitch(ep))
	fmt.Println(dynamicSwitch("string"))
}
