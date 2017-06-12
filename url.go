package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func dogoquery() {
	doc, err := goquery.NewDocument("http://endoyuta.com")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
	})
}
