package main

import (
	"bufio"
	"fmt"
	"os"
)

//Go 的 map 类似于 Java 语言中的 HashMap ，Python语言中
//的 dict ， Lua 语言中的 table ，通常使用 hash 实现
func main() {
	counts :=make(map[string] int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan()  {
		counts[input.Text()]++
	}
	//fmt.Printf(string(counts))
	for line,n:=range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}

}
