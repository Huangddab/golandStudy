package main

import (
	"fmt"
	"os"
)

// 1 打开和关闭文件

// func main() {
// 	// 只读方式打开当前目录下的main.go文件
// 	file, err := os.Open("./main.go")
// 	defer file.Close() //必须得关闭文件流
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//操作文件
// 	fmt.Println(file) //&{0xc00014a780}
// }

// 2 file.Read() 读取文件

// func main() {
// 	// 只读的方式打开当前目录下的main.go文件
// 	file, err := os.Open("./main.go")
// 	if err != nil {
// 		// 代表有错误
// 		fmt.Println("打开文件失败，错误是：", err)
// 		return
// 	}

// 	// 关闭文件
// 	defer file.Close()

// 	// 使用Read方法读取文件数据，注意一次会读取128个字节
// 	tmp := make([]byte, 128)
// 	n, err := file.Read(tmp)
// 	if err == io.EOF {
// 		// 代表文件读完了
// 		fmt.Println("文件读完了")
// 		return
// 	}
// 	if err != nil {
// 		// 代表有错误
// 		fmt.Println("读取文件失败，错误是：", err)
// 		return
// 	}
// 	fmt.Printf("读取了%d字节数据\n", n)
// 	fmt.Println(string(tmp[:n])) //将读取的字节数据转换为字符串并打印出来。
// }

//  3 循环读取

// func main() {
// 	file, err := os.Open("./main.go")
// 	if err != nil {
// 		fmt.Println("打开文件失败，错误是：", err)
// 	}
// 	defer file.Close()

// 	// 循环读取文件
// 	var content []byte
// 	var tmp = make([]byte, 128)
// 	for {
// 		n, err := file.Read(tmp)
// 		if err == io.EOF {
// 			fmt.Println("文件读完了")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("读取文件失败，错误是：", err)
// 			return
// 		}
// 		content = append(content, tmp[:n]...)
// 	}
// 	fmt.Println(string(content))
// }

// 4 bufio 读取文件

// func main() {
// 	file, err := os.Open("./note.txt")
// 	if err != nil {
// 		fmt.Println("打开文件失败，错误是：", err)
// 		return
// 	}
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	// 按行读取文件
// 	for {
// 		line, err := reader.ReadString('\n')
// 		// reader.ReadString('\n') 从文件中读取一行，直到遇到换行符 \n 或文件结束 (io.EOF)。
// 		if err == io.EOF {
// 			// 代表文件读完了
// 			if len(line) != 0 {
// 				// 代表最后一行没有换行符，需要单独打印出来
// 				fmt.Println(line)
// 			}
// 			fmt.Println("文件读完了")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("读取文件失败，错误是：", err)
// 			return
// 		}
// 		// 打印读取的内容，每一行都有一个换行符
// 		fmt.Print(line)
// 	}
// }

// 5 ioutil 读取整个文件

// func main() {
// 	content, err := ioutil.ReadFile("./note.txt")
// 	if err != nil {
// 		fmt.Println("读取文件失败，错误是：", err)
// 		return
// 	}
// 	fmt.Println(string(content))
// }

// 6 文件写入操作

//  ( 1 ) Write 和 WriteString

// func main() {
// 	// os.O_WRONLY|os.O_APPEND 表示以只写方式打开文件，并将写入的数据追加到文件末尾。
// 	// 0666 表示文件的权限，这里是读写权限。
// 	file, err := os.OpenFile("./note.txt", os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		fmt.Println("打开文件失败，错误是：", err)
// 		return
// 	}
// 	defer file.Close()
// 	str := "你好，世界！"
// 	file.Write([]byte(str))  //将字符串转换为字节切片并写入文件。
// 	file.WriteString("\r\n") //写入一个换行符。
// 	file.WriteString("直接写入字符串数据")
// }

// ( 2 ) bufio.NewWriter

// func main() {
// 	file, err := os.OpenFile("./note.txt", os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		fmt.Println("打开文件失败，错误是：", err)
// 		return
// 	}
// 	defer file.Close()
// 	writer := bufio.NewWriter(file)
// 	// 循环写入10行数据
// 	for i := 0; i < 10; i++ {
// 		writer.WriteString("你好，世界！bufio.NewWrite\r\n") //将字符串写入缓冲区。
// 	}
// 	writer.Flush() //将缓冲区中的数据写入文件。
// }

//  ( 3 ) ioutil.WriteFile

// func main() {
// 	// ioutil.WriteFile 会清空文件内容，然后将新内容写入文件。
// 	str := "你好，世界！ioutil.WriteFile\r\n"
// 	err := ioutil.WriteFile("./note.txt", []byte(str), 0666)
// 	if err != nil {
// 		fmt.Println("写入文件失败，错误是：", err)
// 		return
// 	}
// }

// 7 文件重命名
//
//	func main() {
//		err := os.Rename("./note.txt", "./note2.txt")
//		if err != nil {
//			fmt.Println("重命名失败，错误是：", err)
//			return
//		}
//	}
// func main() {
// 	err := os.Rename("./note2.txt", "./note.txt")
// 	if err != nil {
// 		fmt.Println("重命名失败，错误是：", err)
// 		return
// 	}
// }

// 8 复制文件

// (1) 第一种复制文件方法：ioutil 进行复制

// // 接收两个文件路径 srcFileName dstFileName
// func CopyFile(dstFileName string, srcFileName string) (err error) {
// 	// 读取源文件
// 	input, err := ioutil.ReadFile(srcFileName)
// 	if err != nil {
// 		fmt.Println("读取文件失败，错误是：", err)
// 		return
// 	}
// 	// 将读取的内容写入目标文件
// 	err = ioutil.WriteFile(dstFileName, input, 0666)
// 	if err != nil {
// 		fmt.Println("写入文件失败，错误是：", dstFileName, err)
// 		return
// 	}
// 	return nil
// }
// func main() {
// 	srcFileName := "./note.txt"
// 	dstFileName := "./copy/note3.txt"
// 	// 调用 CopyFile 函数进行文件复制
// 	err := CopyFile(dstFileName, srcFileName)
// 	if err == nil {
// 		fmt.Println("复制文件成功")
// 	} else {
// 		fmt.Println("复制文件失败，错误是：", err)
// 	}
// }

// (2) 复制文件方法流

// 接收两个文件路径 srcFileName dstFileName
// func CopyFile(dstFileName string, srcFileName string) (err error) {
// 	// 打开源文件
// 	source, _ := os.Open(srcFileName)
// 	// 打开目标文件
// 	destination, _ := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)

// 	buf := make([]byte, 128)
// 	for {
// 		// 从源文件读取数据
// 		n, err := source.Read(buf)
// 		if err != nil && err != io.EOF {
// 			// 读取文件失败，并且不是文件结束的错误
// 			return err
// 		}
// 		// 读取结束，退出循环
// 		if n == 0 {
// 			break
// 		}
// 		// 将读取的数据写入目标文件 buf[:n] 表示写入 buf 切片中的前 n 个字节数据。
// 		if _, err := destination.Write(buf[:n]); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	// 调用CopyFile函数进行文件复制
// 	srcFile := "./note.txt"
// 	dstFile := "./copy/note4.txt"
// 	err := CopyFile(dstFile, srcFile)
// 	if err == nil {
// 		fmt.Println("复制文件成功")
// 	} else {
// 		fmt.Println("复制文件失败，错误是：", err)
// 	}
// }

// 9 创建目录

// (1) 一次创建一个目录

// func main() {
// 	err := os.Mkdir("./test", 0666)
// 	if err != nil {
// 		fmt.Println("创建目录失败，错误是：", err)
// 	}
// }

// (2) 一次创建多级目录

// func main() {
// 	err := os.MkdirAll("./test/a/b/c", 0666)
// 	if err != nil {
// 		fmt.Println("创建目录失败，错误是：", err)
// 	}
// }

// 10 删除目录和文件

// (1) 删除一个目录或文件
// func main() {
// 	// 删除目录
// 	// err := os.Remove("./test/a/b/c")
// 	// 删除文件
// 	err := os.Remove("./copy/note3.txt")
// 	if err != nil {
// 		fmt.Println("删除目录失败，错误是：", err)
// 	}
// }

// (2) 删除多级目录或文件
func main() {
	// 删除多级目录及文件
	// err := os.RemoveAll("./test")
	err := os.RemoveAll("./copy")
	if err != nil {
		fmt.Println("删除目录失败，错误是：", err)
	}
}
