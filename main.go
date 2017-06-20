package main

import (
	"fmt"
	"log"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db := db_init()
	defer func(){
		fmt.Println("Closing DB...")
		db.Close()
	}()

	val, err := db.Get([]byte("key"), nil)
	if err != nil {
		fmt.Println("Key is not exist")
		put(db, []byte("key"), []byte("hello"))
		val, err = db.Get([]byte("key"), nil)
		if err != nil {
			log.Fatal("get error")
		}
	}
	fmt.Printf("key is %s\n", string(val))

	del(db, []byte("key"))
}

func db_init() *leveldb.DB {
	fmt.Println("open database")
	db, err := leveldb.OpenFile("./data/db", nil)
	if err != nil {
		log.Fatal("open error")
	}
	return db
}

func put(db *leveldb.DB, key []byte, value []byte) {
	fmt.Println("put to db")
	err := db.Put(key, value, nil)
	if err != nil {
		log.Fatal("put error")
	}
}

func del(db *leveldb.DB, key []byte) {
	fmt.Println("del from db")
	err := db.Delete(key, nil)
	if err != nil {
		log.Fatal("delete error")
	}
}
