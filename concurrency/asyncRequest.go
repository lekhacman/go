package concurrency

import (
	"net/http"
	"log"
	"io/ioutil"
)

func asyncGet(url string, c chan string) {
	getResp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer getResp.Body.Close()
	respB, err := ioutil.ReadAll(getResp.Body)

	c <- string(respB)
}

func GetPingPong() string {
	var pingPong string

	c := make(chan string)

	go asyncGet("http://localhost:8000/ping", c)
	go asyncGet("http://localhost:8000/pong", c)

	r1, r2 := <-c, <-c

	pingPong = r1 + r2

	return pingPong
}