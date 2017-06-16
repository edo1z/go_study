package main

import (
	"fmt"
	"net/http"
	"log"
	"golang.org/x/net/context/ctxhttp"
	"context"
	"time"
)

var url = "https://www.imagemagick.org/download/linux/CentOS/x86_64/ImageMagick-7.0.6-0.x86_64.rpm"
var out_dir = "~/"

func main() {
	fmt.Println("test file download")
	check()
	download()
}

func check() {
	fmt.Println("checking url...")
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	res, err := ctxhttp.Head(ctx, http.DefaultClient, url)
	chkErr(err)
	if res.Header.Get("Accept-Ranges") != "bytes" {
		log.Fatal("not supported range access.")
	}
	if res.ContentLength <= 0 {
		log.Fatal("invalid content length.")
	}
	fmt.Println("check ok.")
}

func download(){
	fmt.Println(out_dir)
	ProgressBar()
}

func chkErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}