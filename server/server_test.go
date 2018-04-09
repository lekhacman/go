package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"os"
)

func TestPing(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8000/ping", nil)
	w := httptest.NewRecorder()

	ping(w, req)
	assert.Equal(t, 200, w.Code)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, "pong", string(body[:]))
}

func TestPong(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8000/pong", nil)
	w := httptest.NewRecorder()

	pong(w, req)
	assert.Equal(t, 200, w.Code)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, "ping", string(body[:]))
}

//func TestMain(t *testing.T)  {
//	main()
//	var res io.Reader
//	httptest.NewRequest("GET", "/ping", res)
//	assert.Equal(t, "server", "server")
//}

func TestMain(m *testing.M)  {
	defer os.Exit(m.Run())
}