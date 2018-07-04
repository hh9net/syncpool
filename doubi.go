package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func getBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

func putBuffer(buf *bytes.Buffer) {
	if buf.Len() > 1024 {
		return
	}
	buf.Reset()
	bufPool.Put(buf)
}

func main() {
	buf := getBuffer()

	buf.WriteString("=?")
	buf.WriteByte('?')
	buf.WriteByte('q')
	buf.WriteByte('?')

	buf.WriteString("?=")

	es := buf.String()
	putBuffer(buf)
	fmt.Println(es)
	uu := getBuffer()
	fmt.Println(uu.String())
}
