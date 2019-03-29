package main

import (
	"os"
	"strings"
)

func writerTo() {
	reader := strings.NewReader("golang中文网")
	// reader := bytes.NewReader([]byte("golang中文网"))
	reader.WriteTo(os.Stdout) // func (*Reader) WriteTo	func (b *Reader) WriteTo(w io.Writer) (n int64, err error)	WriteTo实现了io.WriterTo。
}

func main() {
	writerTo()
}
