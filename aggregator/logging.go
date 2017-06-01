package aggregator

import (
	"bytes"
	"os"
	"sync"
	"time"
)

type Logging struct {
	file *os.File
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func NewLogging(file *os.File) *Logging {
	return &Logging{file}
}

func (self *Logging) Log(title, link string) {
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

func (self *Logging) Close() error {
	return self.file.Close()
}
