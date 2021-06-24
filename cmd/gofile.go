package main

import (
	"fmt"
	"io"
	"os"
	"flag"
	"bufio"
	"runtime"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func appendToFile(f *os.File, content string) error {
	// 以只写的模式，打开文件

	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	_, err := f.WriteAt([]byte(content), n)

	return err
}

func main() {
	filePath := flag.String("f", "mm.txt", "Input filePath:")
	flag.Parse()


	saveContentFile, err := os.Create("ok.txt")
	if err != nil {
		fmt.Println("Save Content file create failed. err: " + err.Error())
	}

	fmt.Printf(*filePath)

	FileHandle, err := os.Open(*filePath)
		if err != nil {
		panic("Read dir error, [" + err.Error() + "]")
	}
	defer FileHandle.Close()
	lineReader := bufio.NewReader(FileHandle)
	readCurrent := func(ch *chan string) {
		defer close(*ch)
		line, _, err := lineReader.ReadLine()
		if err != io.EOF {
			*ch <- string(line)
		} else {
			*ch <- "end"
		}

	}

	for {
		ch := make(chan string, 1000)
		go readCurrent(&ch)

		for j := 0; j < 1000; j++ {
			content, err := <-ch
			if content == "end" {
				runtime.Goexit()
				break
			}
			if err != false {
				appendToFile(saveContentFile, content+"\n")
			}
		}

	}

	defer saveContentFile.Close()

}