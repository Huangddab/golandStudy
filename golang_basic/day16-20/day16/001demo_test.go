package day16

import (
	"fmt"
	"os"
	"testing"
)

func TestDemo1(t *testing.T) {
	fileInfo, err := os.Stat("./doc/aa.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("fileInfo:", fileInfo)
	fmt.Println("filename:", fileInfo.Name())
	fmt.Println("size:", fileInfo.Size())
	fmt.Println("isDir:", fileInfo.IsDir())
	fmt.Println("modTime:", fileInfo.ModTime())
	fmt.Println("mode:", fileInfo.Mode())
	fmt.Println("name:", fileInfo.Name())
	fmt.Println("size:", fileInfo.Size())
	fmt.Println("sys:", fileInfo.Sys())
}
