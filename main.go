package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	data, err := Assets.Open("/resessh.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	b := new(bytes.Buffer)
	io.Copy(b, data)
	buf := b.Bytes()

	fmt.Println(string(buf))
}
