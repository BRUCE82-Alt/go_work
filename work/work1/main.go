package main

import "fmt"

var ch = make(chan int)

func Print(a string) {
	fmt.Println(a)
	ch <- 0
}

func main() {
	for i := 0; i <= 10; i++ {
		go Print("张三")
		<-ch
		go Print("李四")
		<-ch
		go Print("王五")
		<-ch
		go Print("赵六")
		<-ch
	}

}
