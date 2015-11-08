package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	tr := &http.Transport{}
	tr.RegisterProtocol("file", http.NewFileTransport(http.Dir("/"))) // http.Dir()は定義したProtocolのrootディレクトリを指定
	c := &http.Client{Transport: tr}
	r, _ := c.Get("file:///.DS_Store") // rootディレクトリの.DS_Storeにクライアントからアクセスする
	io.Copy(os.Stdout, r.Body)         // Copy(dst,src)なのでGetした結果(r.Body)を標準出力(os.Stdout)に表示する
}
