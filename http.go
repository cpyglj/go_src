package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var Count = 1

//生成随机字符串
func GetRandomString() string {
	Count += 1
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < Count; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	time.Sleep(4 * time.Second)
	io.WriteString(w, "请求id:"+id+"\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
