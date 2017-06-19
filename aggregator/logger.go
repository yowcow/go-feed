package aggregator

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	file *os.File
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func NewLogger(file *os.File) *Logger {
	return &Logger{file}
}

func (self *Logger) Log(title, link string) {
	b := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(b)

	b.Reset()
	b.WriteString(time.Now().Format(time.StampNano))
	b.WriteByte(' ')
	b.WriteString(link)
	b.WriteByte(' ')
	b.WriteString(title)
	b.WriteString("\n")

	self.file.Write(b.Bytes())
}

func (self *Logger) Close() error {
	return self.file.Close()
}

type LoggerQueue struct {
	Wg  *sync.WaitGroup
	In  chan RssItem
	Out *Logger
}

func LoggerWorker(id int, q LoggerQueue) {
	name := fmt.Sprintf("[Logger Worker %d]", id)
	defer func() {
		fmt.Println(name, "Exiting")
		q.Wg.Done()
	}()
	for item := range q.In {
		q.Out.Log(item.Title, item.Link)
	}
}
