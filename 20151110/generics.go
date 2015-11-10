package main

import (
	"fmt"
)

type Any interface{}

type GetValuer interface {
	GetValue() Any
}

type Value struct {
	v Any
}

func (v *Value) GetValue() Any {
	return v.v
}

func main() {
	//いろんな型の引数が取れる
	var i GetValuer = &Value{10}
	var s GetValuer = &Value{"vvv"}

	var values []GetValuer = []GetValuer{i, s}

	for _, val := range values {
		fmt.Println(val.GetValue())
	}
}
