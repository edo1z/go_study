package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//コマンドライン引数からファイルパスを取得
func GetFileName() string {
	if len(os.Args) < 2 {
		log.Fatal("Please input file name")
	}
	return os.Args[1]
}

//ファイル全体の文字列を返す
func ReadAllOfFile(file_path string) string {
	bs, err := ioutil.ReadFile(file_path)
	ChkErr(err)
	return string(bs)
}

//ファイル全体の文字列を返す
func ReadAllOfFile2(file_path string) string {
	fp, err := os.Open(file_path)
	ChkErr(err)
	bs, err := ioutil.ReadAll(fp)
	ChkErr(err)
	return string(bs)
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

//ファイル全体を表示
func ReadFile(file_path string) {
	fp, err := os.Open(file_path)
	ChkErr(err)
	//defer fp.Close()
	reader := bufio.NewReader(fp)
	str, err := Readln(reader)
	for err == nil {
		fmt.Println(str)
		str, err = Readln(reader)
	}
}

func WriteFile() string {
	return "string"
}

//func OpenOrCreateFile() *os.File {

//}

func ChkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
