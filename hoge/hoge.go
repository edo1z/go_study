package hoge

import (
	"path/filepath"
	"fmt"
	"os"
)

var basepath = "./data/"

func Hoge(f string) {
	fpath := filepath.Join(basepath, f)
	info, err := os.Stat(fpath)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(info.Name())
		fmt.Println(info.Size())
		fmt.Println(info.Mode())
		fmt.Println(info.ModTime())
		fmt.Println(info.IsDir())
		fmt.Println(info.Sys())
	}
}
