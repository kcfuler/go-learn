package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downStream chan string) {
	for _, v := range []string{"hello world", "bad string", "goodBye !"} {
		downStream <- v
	}
	close(downStream) // 每一个通道的打开都要对应一个关闭
}

func filterGopher(upstream, downSteam chan string) {
	for item := range upstream { // 读取通道中的值，直到通道关闭
		if !strings.Contains(item, "bad") {
			downSteam <- item
		}
	}
	close(downSteam)
}

func printGopher(upStream chan string) {
	for v := range upStream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string) // 通过 make 创建一个通道
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
	// strs :=
}
