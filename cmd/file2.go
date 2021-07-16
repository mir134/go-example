package main

import (
	"github.com/mir134/go-example/pkg/file"
	"flag"
	"fmt"
)
func main() {
	filePath := flag.String("f", "mm.txt", "处理文件路经:")
	ssfnPath := flag.String("s", "ssfn6759621566561569682", "SSFN文件路经:")
	fileType := flag.String("t", "2", "处理类型:")
	flag.Parse()
	fmt.Println("处理：" + *filePath , ", 处理类型: " + *fileType +", 复制ssfn："+ *ssfnPath + "")
	file.ReadEachLineReaderCreateDir(*filePath, *ssfnPath, *fileType)
	//file.MkUserDir("ddsfdsfdsf", "23323123" , *ssfnPath)
}
