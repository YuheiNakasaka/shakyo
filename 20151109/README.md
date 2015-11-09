# Goでフォーム入力

フォーム関係ないけどinterfaceについて理解深まった

## 写経元

- [https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/04.2.md](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/04.2.md)

# 学び

- 写経元の著者の[独自ライブラリ](https://github.com/astaxie/beeku/blob/master/slice.go)にある関数Slice_diffとIn_sliceをserver.goに移すとこで少しハマった
- 元ライブラリでは引数に[]interface{}型になっていたが実際は[]stringであり、エラーが吐かれて迷った
- 調べてみると、interface{}及び[]interface{}型を引数に取るときは特定の型を取らないときに使うらしい
- []interface{}型に[]stringを代入できなかったのは[]stringの中身をinterfaceに変換できないから([https://golang.org/doc/faq#convert_slice_of_interface](https://golang.org/doc/faq#convert_slice_of_interface))
- よって、[]stringを[]inteface{}に代入するには以下のようにforを回して変換する必要がある
```
func f(slice []interface{}) (rslice []interface{}) {
  for i,v := range  slice {
    rslice[i] = v
  }
  return rslice
}

slice := []string{"football", "basketball", "tennis"}
converted_slice := make([]interface{}, len(slice))
for i, v := range slice {
  converted_slice[i] = v
}

f(slice) // cannot convert
f(converted_slice) // OK
```

- 参考
  - [Go言語の型とreflect](http://qiita.com/atsaki/items/3554f5a0609c59a3e10d)
  - [Go の interface 設計](http://jxck.hatenablog.com/entry/20130325/1364251563)