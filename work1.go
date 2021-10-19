package main

import (
	"fmt"
	"time"
)
//张三, 李四,王五,赵六10次
func main() {
	channel1,channel2,channel3,channel4:=make(chan int),make(chan int),make(chan int),make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case <-channel4:
				fmt.Println("张三")
			channel1<-i
			}
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case <-channel1:
				fmt.Println("李四")
			}
			channel2<-i
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {

			select {
			case <-channel2:
				fmt.Println("王五")
			}
			channel3<-i
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			if i == 0 {
				channel4<-0
			}
			select {
			case <-channel3:
				fmt.Println("赵六")
				channel4<-i
			}
		}
	}()
	time.Sleep(2*time.Second)
}