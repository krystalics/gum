package concurrency

import (
	"errors"
	"fmt"
	"time"
)

// @Author linjiabao
// @Date   2022/5/22
// 存放公共的struct与方法

// SyncGet 执行fn、可设置超时时间
func SyncGet[T any](fn func() T, second int) (T, error) {
	timeout := time.After(time.Duration(second) * time.Second)
	c := make(chan T)
	go func() {
		c <- fn()
	}()

	select {
	case <-timeout:
		//超时
		var zero T
		return zero, errors.New("AsyncGet timeout")
	case value := <-c:
		return value, nil
	}
}

func WriteResultAfter(ch chan int, second int) {
	timeout := time.After(time.Duration(second) * time.Second)

	select {
	case <-timeout:
		fmt.Println("it's the time")
		ch <- 1
	}
}
