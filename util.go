package main

import (
	"gopkg.in/cheggaaa/pb.v2"
	"time"
	"net/url"
	"path/filepath"
	"os"
)

func GetFileName(u string) string {
	result, err := url.Parse(u)
	chkErr(err)
	return filepath.Base(result.Path)
}

func GetFilePath(d, f string) string {
	return filepath.Join(d, f)
}

func ProgressBar(t *Target, file_path *string, ch <-chan string){
	file_size := t.file_size
	now_size := int64(0)
	bar := pb.Start64(file_size)
	for {
		select {
		case <-ch:
			bar.SetCurrent(file_size)
			bar.Finish()
			break
		default:
			fi, err := os.Stat(*file_path)
			chkErr(err)
			now_size = fi.Size()
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
}