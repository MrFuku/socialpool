package main

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/garyburd/go-oauth/oauth"
)

var conn net.Conn

func dial(netw, addr string) (net.Conn, error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}
	netc, err := net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil {
		return nil, err
	}
	conn = netc
	return conn, nil
}

var reader io.ReadCloser

func closeConn() {
	if conn != nil {
		conn.Close()
	}
	if reader != nil {
		reader.Close()
	}
}

var (
	authClient *oauth.Client
	creds      *oauth.Credentials
)

func setupTwitterAuth() {
	creds = &oauth.Credentials{
		Token:  os.Getenv("SP_TWITTER_ACCESSTOKEN"),
		Secret: os.Getenv("SP_TWITTER_ACCESSSECRET"),
	}
	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  os.Getenv("SP_TWITTER_KEY"),
			Secret: os.Getenv("SP_TWITTER_SECRET"),
		},
	}
}
