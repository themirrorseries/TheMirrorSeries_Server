package Global

/*
import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

type logFileWriter struct {
	file *os.File
	//write count
	size int64
}

func init() {
	file, err := os.OpenFile("./Logs/mylog"+strconv.FormatInt(time.Now().Unix(), 10), os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
	if err != nil {
		log.Fatal("log  init failed")
	}

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileWriter := logFileWriter{file, info.Size()}
	log.SetOutput(&fileWriter)
	log.Info("start.....")
}
func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}
	n, e := p.file.Write(data)
	p.size += int64(n)
	//文件最大 64K byte
	if p.size > 1024*64 {
		p.file.Close()
		fmt.Println("log file full")
		p.file, _ = os.OpenFile("./Logs/mylog"+strconv.FormatInt(time.Now().Unix(), 10), os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
		p.size = 0
	}
	return n, e
}
*/
