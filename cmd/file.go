package main

import (
	"github.com/mir134/go-example/pkg/file"
	"flag"
	"fmt"
)
func main() {
	filePath := flag.String("f", "mm.txt", "Input filePath:")
	flag.Parse()
	fmt.Println(*filePath)
	file.ReadEachLineReader(*filePath)
}
