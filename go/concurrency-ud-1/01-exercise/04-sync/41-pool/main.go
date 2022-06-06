package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

//TODO: create pool of bytes.Buffers which can be reused.
var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocating new bytes.buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	bufferPool := bufPool.Get().(*bytes.Buffer)
	bufferPool.Reset()

	bufferPool.WriteString(time.Now().Format("15:04:05"))
	bufferPool.WriteString(" : ")
	bufferPool.WriteString(val)
	bufferPool.WriteString("\n")

	w.Write(bufferPool.Bytes())
	bufPool.Put(bufferPool)
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
