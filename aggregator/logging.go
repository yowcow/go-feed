package aggregator

import (
	"bytes"
	"os"
	"sync"
	"time"
)

type Log struct {
	file *os.File
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func NewLogger(file *os.File) *Log {
	return &Log{file}
}

func (self *Log) Log(title, link string) {
	b := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(b)

	b.Reset()
	b.WriteString(time.Now().Format(time.StampNano))
	b.WriteByte(' ')
	b.WriteString(title)
	b.WriteByte(' ')
	b.WriteString(link)
	b.WriteString("\n")

	self.file.Write(b.Bytes())
}

func (self *Log) Close() error {
	return self.file.Close()
}
