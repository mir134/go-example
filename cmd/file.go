package main

import (
	"github.com/mir134/go-example/pkg/file"
	"flag"
	"fmt"
)

func main() {
	dic :=  map[string]string{
		"1": "密码首字母大写",
		"2": "替换用户名数字",
		"3": "用户名前加字符 k (-s k)",
		"4": "用户名密码互换",
		"5": "密码大写",
		"6": "替换密码数字",
		"7": "替换用户名截取少最后一位",
		"8": "密码后加字符1( -s 1)",
		"9": " 账号后加字符1 (-s 1)",
		"10": "密码前加字符叹号 (-s !)",
		"11": "替换账号@后面字符",
		"12": "替换用户名截取少第一位",
	}
	filePath := flag.String("f", "mm.txt", "处理文件路经:")
	fileType := flag.String("t", "11", "处理类型:")
	addStr := flag.String("s", "k", "添加字符:")
	flag.Parse()
	fmt.Println("处理：" + *filePath , "，处理类型: " + dic[*fileType] +"(添加字符: "+ *addStr+ ")，生成：ok_"+ *fileType + ".txt")
	file.ReadEachLineReader(*filePath, *fileType, *addStr)
}
