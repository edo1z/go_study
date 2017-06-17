package main

import (
	"fmt"
	"net/http"
	"log"
	"golang.org/x/net/context/ctxhttp"
	"context"
	"time"
	"os"
	"io"
)

var target_url = "https://www.imagemagick.org/download/linux/CentOS/x86_64/ImageMagick-7.0.6-0.x86_64.rpm"
var out_dir = "."

type Target struct {
	url string
	dir_path string
	file_name string
	file_size int64
}

func main() {
	fmt.Println("test file download")
	target := Target{
		url: target_url,
		dir_path: out_dir,
	}
	check(&target)
	download(&target)
}

func check(t *Target) {
	fmt.Println("checking url...")
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	res, err := ctxhttp.Head(ctx, http.DefaultClient, t.url)
	chkErr(err)
	if res.Header.Get("Accept-Ranges") != "bytes" {
		log.Fatal("not supported range access.")
	}
	if res.ContentLength <= 0 {
		log.Fatal("invalid content length.")
	}
	t.file_size = int64(res.ContentLength)
	t.file_name = GetFileName(t.url)
	fmt.Println("check ok.")
}

func download(t *Target){
	fmt.Println("downloading...")
	req, err := http.NewRequest("GET", t.url, nil)
	chkErr(err)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", 0, t.file_size))
	res, err := http.DefaultClient.Do(req)
	chkErr(err)
	defer res.Body.Close()

	out_path := GetFilePath(t.dir_path, t.file_name)
	out, err := os.Create(out_path)
	chkErr(err)
	defer out.Close()
	ch := make(chan string)
	go ProgressBar(t, &out_path, ch)
	io.Copy(out, res.Body)
	ch <- "fin"
}

func chkErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}