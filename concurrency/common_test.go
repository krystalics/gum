package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// @Author linjiabao
// @Date   2022/5/22

func TestTimeout(t *testing.T) {
	f := func() int {

		return 1
	}

	value, err := SyncGet[int](f, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("value get = ", value)

	time.Sleep(30 * time.Second)
}
