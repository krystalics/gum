package util

import (
	"fmt"
	"testing"
)

// @Author linjiabao
// @Date   2022/5/24

func TestTransfer(t *testing.T) {
	in, out := make(chan int), make(chan int)
	go transfer(in, out)
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
	for i := range out {
		fmt.Println(i)
	}
}
