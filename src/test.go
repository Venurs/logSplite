package main

import (
	"fmt"
	"logSplite/common"
)

func main() {
	filePool := &common.FilePool{
		Length:0,
	}
	filePool.GetFileObj("/Users/manasseh/file_test/2019/03/07/action.log")
	filePool.GetFileObj("/Users/manasseh/file_test/2019/03/06/action.log")
	filePool.GetFileObj("/Users/manasseh/file_test/2019/01/01/action.log")
	fmt.Println(filePool.Length)
	for _, obj := range filePool.FilePool{
		fmt.Println(obj)
	}
	fmt.Println("-------------------------------------------------------------")
	filePool.GetFileObj("/Users/manasseh/file_test/2019/03/06/action.log")
	for _, obj := range filePool.FilePool{
		fmt.Println(obj)
	}
	fmt.Println("-------------------------------------------------------------")
	filePool.GetFileObj("/Users/manasseh/file_test/2019/03/07/action.log")
	for _, obj := range filePool.FilePool{
		fmt.Println(obj)
	}
	filePool.CloseFilePool()
}
