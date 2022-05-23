package bug

import "sync"

// @Author linjiabao
// @Date   2022/5/22

//类似于java的future
type future[T any] interface {
	Get() T //wait for result
}

//第一次设计写的垃圾代码、存在deadlock的bug
func Future[T any](f func() chan T, exec func(t T)) func() {
	c := make(chan T)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c = f()
	}()

	go func() {
		defer wg.Done()
		//在执行完f()之前 c 并不是f()里的channel、所以就会造成sender和receiver不匹配、deadlock
		for t := range c {
			exec(t)
		}
	}()

	//有信息从channel返回
	return wg.Wait
}
