package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

var fatalErr error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

func main() {
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()

	log.Println("データベースに接続します...")
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"mongo"},
		Timeout:  10 * time.Second,
		Username: "root",
		Password: "example",
	}
	db, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		fatal(err)
		return
	}
	defer func() {
		log.Println("データベース接続を閉じます...")
		db.Close()
	}()
	pollData := db.DB("ballots").C("polls")
}
