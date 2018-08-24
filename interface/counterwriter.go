package main

import (
	"bytes"
	"fmt"
	"io"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	r := int64(len(fmt.Sprintf("%s", w)))
	return w, &r
}

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello world"))
	_, v := CountingWriter(&b)
	fmt.Println(*v)
}
