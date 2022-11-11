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
	file	string //目标文件
	key	string //匹配关键字
	dir	string //临时文件
)

func tmpfileWrite(tmpfilePath string,tmpString string) {
	f, err := os.OpenFile(tmpfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	// 关闭文件
	defer f.Close()
	// 字节方式写入
	_, err = f.Write([]byte(tmpString+"\n"))
	if err != nil {
		log.Println(err)
		return
	}
}
// 读取文件的每一行
func readEachLineReader(filePath string,tmpfilePath string, keyString string) {
	start1 := time.Now()
	FileHandle, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer FileHandle.Close()
	lineReader := bufio.NewReader(FileHandle)
	for {
		// 相同使用场景下可以采用的方法
		// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		// func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
		// func (b *Reader) ReadString(delim byte) (line string, err error)
		line, _, err := lineReader.ReadLine()
		res := strings.Contains(string(line), keyString)
		if res{	
			fmt.Println(`del:`+string(line))			
		}else{
			tmpfileWrite(tmpfilePath,string(line))
		}
		if err == io.EOF {
			break
		}
		//fmt.Println(string(line))
	}
	fmt.Println("spend : ", time.Now().Sub(start1))
}
func writeEachLinewriter(filePath string) {

}
func main(){
	flag.StringVar(&file, "file", "", "log file path")
	flag.StringVar(&key, "key", "", "keywords to match")
	flag.Parse()
	dir,_ := os.Getwd()
	var tempfile string = string(dir)+`/temp.tmp`
	readEachLineReader(file,tempfile,key)
	fileInfo, err := os.Stat(tempfile)
	if err != nil {
		fmt.Println("Temporary file creation failed,can't overwrite!")
		fmt.Println("**********Please enter -h for help!************")
		log.Println(err)
		return
	}
	
	if len(fileInfo.Name())>0{
		os.Rename(tempfile,file)
	}
}
