package file

import (
"fmt"
"io"
"log"
"os"
"time"
	"bufio"
	"strings"
	"regexp"
	"path/filepath"
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

func ReadEachLineReader(filePath string, fileType string, addStr string) {
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
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], addStr + countSplit[0], 1)) // 用户名前加q
				case "4":
					fmt.Fprintln(f,  countSplit[1] + "----" + countSplit[0]) // 用户名密码互换
				case "5":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], UpperCase(countSplit[1]), 1)) // 密码大写
				case "6":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], ReplaceNumber(countSplit[1]), 1)) // 替换密码数字
				case "7":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], countSplit[0][0 : len(countSplit[0])-1], 1)) // 替换用户名截取少最后一位
				case "8":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], countSplit[1] + addStr, 1)) // 密码加1
				case "9":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], countSplit[0]+ addStr, 1)) // 账号后加1
				case "10":
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], addStr + countSplit[1] , 1)) // 密码前叹号1
				case "11":
					pat := "@\\S+$"
					re, _ := regexp.Compile(pat)
					str := re.ReplaceAllString(countSplit[0], "")
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], str, 1)) // 替换账号@后面字符
				case "12":
					if len(countSplit[0]) > 1 {
						fmt.Fprintln(f, strings.Replace(string(line), countSplit[0], countSplit[0][1:len(countSplit[0])], 1)) // 替换用户名截取少第一位
					}
				default:
					fmt.Fprintln(f, strings.Replace(string(line), countSplit[1], Capitalize(countSplit[1]), 1))  //Capitalize 大写
			}
		}
	}
	fmt.Println("处理用时: ", time.Now().Sub(start1))
}

func ReadEachLineReaderCreateDir(filePath string, ssfnPath string, fileType string) {
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
					MkUserDir(countSplit[0], countSplit[1], ssfnPath)  // 创建文件夹
				case "2":
					//fmt.Println("处理用户: ", countSplit[0])
					CopyUser(countSplit[0]) // 迁移old_data 用户文件夹到 new_data
				default:

			}
		}
	}
	fmt.Println("处理用时: ", time.Now().Sub(start1))
}

func CopyUser(user string)  {
	b, _ := PathExists("old_data\\" + user)
	//fmt.Println("CopyUser文件: ", b)
	if b == true {
		//fmt.Println("文件: ", b)
		b, _ := PathExists("new_data\\")
		if b == false {
			err := os.Mkdir("new_data\\", os.ModePerm) //在当前目录下生成md目录
			if err != nil {
				fmt.Println(err)
			}
		}
		c, _ := PathExists("new_data\\" + user)
		if c == false {
			err := os.Mkdir("new_data\\" + user, os.ModePerm) //在当前目录下生成md目录
			if err != nil {
				fmt.Println(err)
			}
		}
		copyDir("old_data\\" + user, "new_data\\" + user)
		os.RemoveAll("old_data\\" + user)
	}

}

func MkUserDir(user string, pass string, srcFile string) bool{
	os.MkdirAll("ok_data/" + user, os.ModePerm)
	f, err := os.Create("ok_data/"+ user + "/"+ user+ ".txt")
	targFile := "ok_data/"+ user + "/"+ srcFile
	if err != nil {
		log.Println(err)
		return false
	}
	defer f.Close()
	fmt.Fprintln(f,  user + "----" + pass) // 用户名密码互换

	srcSources, err := os.Open(srcFile)
	//if file to open the file err isnt nil
	//but in php before do it we must to use file_exists to check the file is exsits
	if err != nil {
		fmt.Println("cant to open the file check it is exists or not")
		return false
	}
	// create a new file
	tar, err := os.Create(targFile)
	if err != nil {
		fmt.Println("cant create the file")
		return false
	}
	defer tar.Close()
	defer srcSources.Close()
	scanner := bufio.NewScanner(srcSources)
	for scanner.Scan() {
		tar.Write(scanner.Bytes())
		tar.Write([]byte("\n"))
	}
	return true
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

func copyDir(src string, dest string) {
	src_original := src
	err := filepath.Walk(src, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			//fmt.Println(f.Name())
			//os.Mkdir(dest +"\\" + f.Name(), os.ModePerm)
			//fmt.Println("CopyFile:" + src + " to " + dest  )
			//copyDir(src , dest + "\\" + f.Name())
		} else {
			//fmt.Println(src)
			//fmt.Println(src_original)
			//fmt.Println(dest)
			dest_new := strings.Replace(src, src_original, dest, -1)
			//fmt.Println(dest_new)
			fmt.Println("CopyFile:" + src + " to " + dest_new)
			CopyFile(src, dest_new)
		}
		//println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//copy file
func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	//fmt.Println("dst:" + dst)
	dst_slices := strings.Split(dst, "\\")
	dst_slices_len := len(dst_slices)
	dest_dir := ""
	for i := 0; i < dst_slices_len-1; i++ {
		dest_dir = dest_dir + dst_slices[i] + "\\"
	}
	//dest_dir := getParentDirectory(dst)
	//fmt.Println("dest_dir:" + dest_dir)
	b, err := PathExists(dest_dir)
	if b == false {
		err := os.Mkdir(dest_dir, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		}
	}
	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
