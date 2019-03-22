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
	wg.Add(2)
	go logP.Read(&wg, &date)
	go logP.WriteFilePool(&wg)

	wg.Wait()
	fmt.Println(time.Now().Sub(startTime))
}
