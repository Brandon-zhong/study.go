package benchmark_demo

import (
	"gitlab.gaeamobile-inc.net/sp2/gaeaspgo/util"
	"gitlab.gaeamobile-inc.net/sp2/gaeaspgo/util/must"
	"gitlab.gaeamobile-inc.net/sp2/gaeaspgo/util/unwind"
	"gitlab.gaeamobile-inc.net/sp2/gaeaspgo/zlog"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func BenchmarkDemo(b *testing.B) {
	str := ""
	for n := 0; n < b.N; n++ {
		str += strconv.Itoa(n)
	}
}

func BenchmarkWriteFile(b *testing.B) {

	var wait sync.WaitGroup
	wait.Add(1)
	var once sync.Once
	go func() {
		wait.Wait()
		time.Sleep(500 * time.Millisecond)
		filePath := "/data/gata/log/" + time.Now().Format("20060102") + "/" + must.String(os.Hostname()) + "_flkajskf" + strconv.Itoa(time.Now().Hour()) + ".log"
		fileWriter = &zlog.FileWriter{File: filePath}
	}()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(wait.Done)
			now := time.Now()
			writeLog(strconv.FormatInt(now.UnixNano(), 10) + " ---- " + now.Format("20060102"))
		}
	})

}

var fileWriter *zlog.FileWriter = &zlog.FileWriter{File: "/data/gata/log/" + time.Now().Format("20060102") + "/" + must.String(os.Hostname()) + "_" + strconv.Itoa(time.Now().Hour()) + ".log"}
var buffer = make([]string, 0, 1000)
var bufferLock sync.Mutex

func writeLog(s string) {
	bufferLock.Lock()
	defer bufferLock.Unlock()
	if len(buffer) < 1000 {
		buffer = append(buffer, s)
		return
	}
	tmpBuf := buffer
	buffer = make([]string, 0, 1000)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zlog.Error("Unexpected panic",
					zap.Any("error", err),
					zap.String("stack", unwind.Stack(1)))
			}
		}()
		filePath := "/data/gata/log/" + time.Now().Format("20060102") + "/" + must.String(os.Hostname()) + "_" + strconv.Itoa(time.Now().Hour()) + ".log"
		if fileWriter.File != filePath {
			fileWriter = &zlog.FileWriter{File: filePath}
		}
		writer := fileWriter
		_, _ = writer.Write(util.StringToBytes(strings.Join(tmpBuf, "\n") + "\n"))
	}()
}
