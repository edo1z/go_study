package main

import (
	"gopkg.in/cheggaaa/pb.v2"
	"time"
)

func ProgressBar(){
	file_size := int64(1024)
	now_size := int64(0)
	bar := pb.Start64(file_size)

	for {
		now_size += 50
		if now_size < file_size {
			bar.SetCurrent(now_size)
		} else {
			bar.SetCurrent(file_size)
			bar.Finish()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}