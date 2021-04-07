package main

import (
	"gopl.io/ch8/thumbnail"
)

func main() {
	files := [] string {
		"/home/yuanjiawei/图片/avatar.jpg",
		"/home/yuanjiawei/图片/avatar（第 1 个复件）.jpg",
		"/home/yuanjiawei/图片/avatar（第 2 个复件）.jpg",
		"/home/yuanjiawei/图片/avatar（第 3 个复件）.jpg",
		"/home/yuanjiawei/图片/avatar（第 4 个复件）.jpg",
		"/home/yuanjiawei/图片/avatar（第 5 个复件）.jpg",
		"/home/yuanjiawei/图片/avatar（第 6 个复件）.jpg",
		"/home/yuanjiawei/图片/avatar（第 7 个复件）.jpg",
		"/home/yuanjiawei/图片/avatar（第 8 个复件）.jpg",
	}
	makeThumbnails3(files)
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}
	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}
