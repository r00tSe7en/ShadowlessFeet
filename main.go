package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"flag"
	"strings"
)

var (
	file string //目标文件
	key string //匹配关键字
	dir string //临时文件
)

// 读取文件的每一行
func readEachLineReader(filePath string,tmpfilePath string, keyString string) {
	var timesNum = 0
	start1 := time.Now()
	FileHandle, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := os.OpenFile(tmpfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	lineReader := bufio.NewReader(FileHandle)
	for {
		line, _, err := lineReader.ReadLine()
		if err == io.EOF {
			break
		}
		res := strings.Contains(string(line), keyString)
		if res{	
			timesNum = timesNum+1	
		}else{
			// 字节方式写入
			_, err = f.Write([]byte(string(line)+"\n"))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
	// 关闭文件
	defer f.Close()
	if timesNum>0{
		os.Rename(tmpfilePath,file)
	}else{
		os.Remove(tmpfilePath)
	}
	//匹配删除的行数
	fmt.Println("delete : ", timesNum)
	fmt.Println("spend : ", time.Now().Sub(start1))
}

func main(){
	flag.StringVar(&file, "file", "", "log file path")
	flag.StringVar(&key, "key", "", "keywords to match")
	flag.Parse()
	dir,_ := os.Getwd()
	var tempfile string = string(dir)+`/temp.tmp`
	if len(file)>0 && len(key)>0{
		readEachLineReader(file,tempfile,key)	
	}else{
		fmt.Println("Please enter -h for help!")
	}
}
