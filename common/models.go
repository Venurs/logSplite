package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type LogProcess struct {
	LogPath     string
	SplitPath   string
	LogChanel   chan string
	SplitRes    chan LogSplitRes
	LOGFilePool FilePool
}

type LogSplitRes struct {
	Date     string
	SplitSli []string
}

func ReadLine(r *bufio.Reader) (string, error) {
	line, isPrefix, err := r.ReadLine()
	for isPrefix && err == nil {
		var bs []byte
		bs, isPrefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename + "/action.log")
	if err == nil || os.IsExist(err) {
		return true
	} else {
		err := os.MkdirAll(filename, os.ModePerm)
		_, err1 := os.Create(filename + "/action.log")
		if err != nil || err1 != nil {
			panic(err1.Error())
		}
		return true
	}
}

func (lp *LogProcess) Read(wg *sync.WaitGroup, date *string) {
	//日志文件读取
	fmt.Println("start read log")
	f, err := os.Open(lp.LogPath)
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		f.Close()

	}()
	r := bufio.NewReader(f)
	for {
		context, err := ReadLine(r)
		if *date != "ALL" && err == nil {
			var logDate string
			if len(context) > 10 {
				logDate = context[:10]
			} else {
				continue
			}
			//logDateParse, _ := time.Parse("2006-01-02", logDate)
			//dateParse, _ := time.Parse("2006-01-02", *date)
			if logDate != *date {
				continue
			}
		}
		if err != nil {
			lp.LogChanel <- "END"
			break
		} else {
			lp.LogChanel <- context
		}
	}
	wg.Done()
}

func (lp *LogProcess) Split(wg *sync.WaitGroup) {
	//日志文件分割
	defer wg.Done()
	logSp := LogSplitRes{}
	var nilSli []string
	for {
		log := <-lp.LogChanel
		if log == "END" {
			break
		}
		var date string
		if len(log) > 10 {
			date = log[:10]
		} else {
			continue
		}

		if logSp.Date != "" {
			if date == logSp.Date {
				logSp.SplitSli = append(logSp.SplitSli, log)
			} else {
				lp.SplitRes <- logSp
				logSp.Date = date
				logSp.SplitSli = nilSli
				logSp.SplitSli = append(logSp.SplitSli, log)
			}
		} else {
			logSp.Date = date
			logSp.SplitSli = append(logSp.SplitSli, log)
		}
	}
	lp.SplitRes <- logSp
	lp.SplitRes <- LogSplitRes{
		Date: "END",
	}
}

func (lp *LogProcess) WriteSlice(wg *sync.WaitGroup) {
	//写入, 每次写入一个切片的日志
	defer wg.Done()
	for {
		logs := <-lp.SplitRes
		if logs.Date == "END" {
			break
		}
		dateSplit := strings.Split(logs.Date, "-")
		filePath := strings.Join(dateSplit, "/")
		filePath = lp.SplitPath + filePath

		if fileExists(filePath) {
			//写入数据
			file, err := os.OpenFile(filePath+"/action.log", os.O_APPEND|os.O_RDWR, os.ModePerm)

			if err != nil {
				panic(err.Error())
			}
			w := bufio.NewWriter(file)
			for _, context := range logs.SplitSli {
				_, err := w.WriteString(context + "\n")
				if err != nil {
					panic(err.Error())
				}
			}
			err = w.Flush()
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(logs.Date + "  write success")
			file.Close()
		}
	}
}

func (lp *LogProcess) WriteFilePool(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		log := <-lp.LogChanel
		if log == "END" {
			break
		} else {
			var date string
			if len(log) > 10 {
				date = log[:10]
			} else {
				continue
			}
			dateSplit := strings.Split(date, "-")
			filePath := strings.Join(dateSplit, "/")
			filePath = lp.SplitPath + filePath
			if fileExists(filePath) {
				file := lp.LOGFilePool.GetFileObj(filePath + "/action.log")
				w := bufio.NewWriter(&file)
				_, err := w.WriteString(log + "\n")
				if err != nil {
					panic(err.Error())
				}
				w.Flush()
			}
		}
	}
	lp.LOGFilePool.CloseFilePool()
}
