package main

import (
	"github.com/mir134/go-example/pkg/file"
	"flag"
	"fmt"
)
func main() {
	filePath := flag.String("f", "mm.txt", "处理文件路经:")
	fileType := flag.String("t", "11", "处理类型:")
	flag.Parse()
	fmt.Println("处理：" + *filePath , ", 处理类型: " + *fileType +", 生成：ok_"+ *fileType + ".txt")
	file.ReadEachLineReader(*filePath, *fileType)
}
