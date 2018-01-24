package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func doGet(client *http.Client, url string, id int) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%d: %s -- %v\n", id, string(buf), err)
	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
func PrintLocalDial(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   3 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}

	fmt.Println("connect done, use", conn.LocalAddr().String())

	return conn, err
}

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

func main() {
	const URL = "http://localhost:8888/"
	client := &http.Client{
		Transport: &http.Transport{
			Dial: PrintLocalDial,
			// DialContext: (&net.Dialer{
			// 	Timeout:   30 * time.Second,
			// 	KeepAlive: 30 * time.Second,
			// }).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
	}

	for {
		go doGet(client, URL+"?id=1", 1)
		go doGet(client, URL+"?id=2", 2)
		go doGet(client, URL+"?id=3", 3)
		time.Sleep(2 * time.Second)
	}
}
