package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	const layout = "20060102"
	t := time.Now()
	dirname := t.Format(layout)
	p, _ := os.Getwd()
	pathname := p + "/" + dirname
	source_file_path := p + "/README.md.template"
	destination_file_path := pathname + "/README.md"

	if _, err := os.Stat(pathname); err != nil {
		if err1 := os.Mkdir(pathname, 0777); err1 != nil {
			fmt.Println("Directory can not be created.")
		}

		src, err := os.Open(source_file_path)
		if err != nil {
			panic(err)
		}
		defer src.Close()

		dst, err := os.Create(destination_file_path)
		if err != nil {
			panic(err)
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			panic(err)
		}

		fmt.Println(dirname + " is successfully created")
	} else {
		fmt.Println(pathname + " is already created")
	}
}
