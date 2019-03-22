package main

import (
	"fmt"
	"logSplite/common"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	logPath, splitPath, date := common.ArgsParse()
	var wg sync.WaitGroup
	logP := &common.LogProcess{
		LogPath:   logPath,
		SplitPath: splitPath,
		LogChanel: make(chan string, 100),
		SplitRes:  make(chan common.LogSplitRes, 100),
	}
	wg.Add(3)
	go logP.Read(&wg, &date)
	go logP.Split(&wg)
	go logP.WriteSlice(&wg)
	wg.Wait()
	fmt.Println(time.Now().Sub(startTime))
}
