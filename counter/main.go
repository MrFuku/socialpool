package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
	"gopkg.in/mgo.v2"
)

var fatalErr error
var countsLock sync.Mutex
var counts map[string]int

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

	log.Println("NSQに接続します...")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())
	if err != nil {
		fatal(err)
		return
	}

	q.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLock.Lock()
		defer countsLock.Unlock()
		if counts == nil {
			counts = make(map[string]int)
		}
		vote := string(m.Body)
		counts[vote]++
		return nil
	}))

	if err := q.ConnectToNSQLookupd("nsqlookupd:4161"); err != nil {
		fatal(err)
		return
	}
}
