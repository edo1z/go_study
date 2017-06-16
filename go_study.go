package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./data/hoge.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(f)
	chkErr(err)
	_ = f.Close()

	thumb := resize.Thumbnail(300, 300, img, resize.Lanczos3)

	out, err := os.Create("./data/thumb.jpg")
	chkErr(err)
	defer Close(out)
	err = jpeg.Encode(out, thumb, nil)
	chkErr(err)
}

func Close(c io.Closer) {
	_ = c.Close()
}

func chkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
