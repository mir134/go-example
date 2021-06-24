package file

import (
"fmt"
"io"
"log"
"os"
"time"
	"bufio"
	"strings"
)

// 文件一块一块的读取
func ReadBlock(filePath string) {
	start1 := time.Now()
	FileHandle, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer FileHandle.Close()
	// 设置每次读取字节数
	buffer := make([]byte, 1024)
	for {
		n, err := FileHandle.Read(buffer)
		// 控制条件,根据实际调整
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if n == 0 {
			break
		}
		// 如下代码打印出每次读取的文件块(字节数)
		fmt.Println(string(buffer[:n]))
	}
	fmt.Println("readBolck spend : ", time.Now().Sub(start1))
}

func ReadEachLineReader(filePath string, fileType string) {
	start1 := time.Now()
	FileHandle, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.Create("ok_"+ fileType+ ".txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer FileHandle.Close()
	defer f.Close()
	lineReader := bufio.NewReader(FileHandle)
	for {
		// 相同使用场景下可以采用的方法
		// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		// func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
		// func (b *Reader) ReadString(delim byte) (line string, err error)
		line, _, err := lineReader.ReadLine()
		if err == io.EOF {
			break
		}
		// 如下是某些业务逻辑操作
		// 如下代码打印每次读取的文件行内容
		countSplit := strings.Split(string(line), "----")
		//fmt.Println(Capitalize(countSplit[1]))
		//fmt.Println(string(line))
		if len(countSplit) > 1 {
			switch fileType {
				case "1":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], Capitalize(countSplit[1]), 1))  // Capitalize 大写
				case "2":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], ReplaceNumber(countSplit[0]), 1)) // 替换用户名数字
				case "3":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], "q" + countSplit[0], 1)) // 用户名加q
				case "4":
					fmt.Fprintln(f,  countSplit[1] + "----" + countSplit[0]) // 用户名密码互换
				case "5":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], UpperCase(countSplit[1]), 1)) // 密码大写
				case "6":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], ReplaceNumber(countSplit[1]), 1)) // 替换密码数字
				case "7":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], countSplit[0][0 : len(countSplit[0])-1], 1)) // 替换用户名截取少一位
				case "8":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], countSplit[1] + "1", 1)) // 密码加1
				case "9":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], countSplit[0]+ "1", 1)) // 账号加1
				case "10":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], countSplit[1] + "!", 1)) // 密码加叹号
				default:
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], Capitalize(countSplit[1]), 1))  //Capitalize 大写
			}
		}
	}
	fmt.Println("处理用时: ", time.Now().Sub(start1))
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				//fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
// Capitalize 字符首字母大写
func ReplaceNumber(str string) string {
	var resStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if vv[i] >= 48 && vv[i] <= 57 {  // 0 到 9

		} else {
			resStr += string(vv[i])
			//fmt.Println("Not begins with lowercase letter,")
			//return str
		}

	}
	return resStr
}

// Upcase 字符字母大写
func UpperCase(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
			vv[i] -= 32 // string的码表相差32位
			upperStr += string(vv[i])
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

