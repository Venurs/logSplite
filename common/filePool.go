package common

import (
	"fmt"
	"os"
	"time"
)

const POOLSIZE int = 10

type FilePoolObj struct {
	Filename string
	File  * os.File
	LatestUse time.Time
}

type FilePool struct {
	Length int
	FilePool [POOLSIZE] FilePoolObj
}

func (fp *FilePool)AddFileObj(filename string) os.File {
	// 添加文件到文件池，如果长度不小于POOLSIZE，则淘汰掉时间最远的一个
	file , _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	filePoolObj := FilePoolObj{
		Filename:filename,
		LatestUse:time.Now(),
		File: file,
	}
	if fp.Length < POOLSIZE {
		fp.FilePool[fp.Length] = filePoolObj
		fp.Length ++
	}else{
		for index, _ := range fp.FilePool{
			if index == fp.Length - 1{
				break
			}
			fp.FilePool[index] = fp.FilePool[index + 1]
			fp.Length ++
		}
	}
	fmt.Println("add file into pool:    " + filename)
	return *filePoolObj.File
}


func (fp *FilePool)GetFileObj(filename string) os.File {
	for index, filePoolObj := range fp.FilePool{
		if filename == filePoolObj.Filename{
			fp.FilePool[index].LatestUse = time.Now()
			var swap = fp.FilePool[index]
			for  i := index; i < fp.Length - 1; i++ {
				fp.FilePool[i] = fp.FilePool[i + 1]
			}

			fp.FilePool[fp.Length - 1] = swap
			return *filePoolObj.File
		}
	}
	return fp.AddFileObj(filename)
}


func (fp *FilePool)CloseFilePool() {
	for _, obj := range fp.FilePool{
		obj.File.Close()
	}
}