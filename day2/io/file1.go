package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("/Users/yangyang/Downloads/abc.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if file != nil {
			err := file.Close()
			if err != nil {
				fmt.Printf("file close error=%v", err)
			}
		}
	}()
	s, err3 := file.WriteString("hello world")
	fmt.Println(s)
	if err3 != nil {
		panic(err3)
	}
	// 读取文件里的内容
	bytes, err2 := os.ReadFile("/Users/yangyang/Downloads/abc.txt")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(bytes))
	n, err := file.WriteAt([]byte("666"), 11)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	// find file last char's index
	length, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
}
