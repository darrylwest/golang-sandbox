package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()

	now := time.Now().UTC()
	b.WriteString(now.Format(time.StampMicro))
	// b.WriteString(string(now.UnixNano()))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	b.WriteByte('\n')
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main() {
	for i := 0; i < 10; i++ {
		Log(os.Stdout, "path", "/search?q=flowers")
		Log(os.Stdout, "xpath", "/search?q=xflowers")
	}

}
