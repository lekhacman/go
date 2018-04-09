package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAsyncGet(t *testing.T) {
	testResponse := "Hello World!"
	testResponse2 := "Hi Man"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testResponse)
	}))

	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testResponse2)
	}))

	defer ts.Close()
	defer ts2.Close()

	c := make(chan string)

	go asyncGet(ts.URL, c)
	go asyncGet(ts2.URL, c)

	res, res2 := <-c, <-c

	assert.Equal(t, testResponse, res)
	assert.Equal(t, testResponse2, res2)

}

func TestGetError(t *testing.T) {
	c := make(chan string)

	go asyncGet("http://localhost:8000", c)
	_, ok := <-c
	assert.False(t, ok)
}
