package concurrency

import (
	"io/ioutil"
	"log"
	"net/http"
)

func asyncGet(url string, c chan string) {
	getResp, err := http.Get(url)
	if err != nil {
		close(c)
		log.Fatal(err)
	}
	defer getResp.Body.Close()
	respB, err := ioutil.ReadAll(getResp.Body)

	c <- string(respB)
}
