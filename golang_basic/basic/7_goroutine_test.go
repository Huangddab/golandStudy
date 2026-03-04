package basic

import (
	"fmt"
	"testing"
)

// 并发任务
func printHello(){
	fmt.Println("Hello")
}

func TestGoroutine(t *testing.T)  {
	go printHello()
	
}