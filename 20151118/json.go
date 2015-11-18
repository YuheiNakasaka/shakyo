package main

import (
	"encoding/json"
	"fmt"
)

type MyType struct {
	A      string
	FooBar string
}

type MyTypeWithTag struct {
	A      string `json:"a"`
	FooBar string `json:"foobar"` // ダブルクォーテーションないと無視される
}

func main() {
	// 構造体をjsonに変換
	mt := MyType{"aaa", "baz"}
	b, _ := json.Marshal(mt)
	fmt.Printf("%s\n", string(b))

	// jsonを構造体に変換
	var mt2 MyType
	json.Unmarshal([]byte(`{"A":"aaa","FooBar":"baz"}`), &mt2)
	fmt.Printf("%#v\n", mt2)

	mt3 := MyTypeWithTag{"aaa", "baz"}
	c, _ := json.Marshal(mt3)
	fmt.Printf("%s\n", string(c))
}
