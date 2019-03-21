package main

import (
	"logSplite/common"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	logP := &common.LogProcess{
		LogPath:   "bb.log",
		SplitPath: "/Users/manasseh/file_test/",
		LogChanel: make(chan string, 100),
		SplitRes:  make(chan common.LogSplitRes, 100),
	}
	wg.Add(2)
	go logP.Read(&wg)
	//go logP.Split(&wg)
	//go logP.WriteSlice(&wg)

	go logP.WriteFilePool(&wg)

	wg.Wait()
}
