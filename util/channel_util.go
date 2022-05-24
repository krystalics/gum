package util

// @Author linjiabao
// @Date   2022/5/24

// transfer provides an unbounded buffer between in and out.  buffer
// exits when in is closed and all items in the buffer have been sent
// to out, at which point it closes out.
func transfer(in <-chan int, out chan<- int) {
	var buf []int
	for in != nil || len(buf) > 0 {
		var i int
		var c chan<- int
		//如果小于0、意味着还没有开始从in写入buffer
		if len(buf) > 0 {
			i = buf[0]
			c = out // enable send case
		}
		//当buf中有数据时、select对应下面两个都可以执行、随机来一个
		select {
		case n, ok := <-in:
			if ok {
				buf = append(buf, n)
			} else {
				in = nil // disable receive case
			}
		case c <- i:
			buf = buf[1:]
		}
	}
	close(out)
}
