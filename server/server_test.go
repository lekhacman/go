package main

import (
	"testing"
	"net/http/httptest"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T)  {
	req := httptest.NewRequest("GET", "http://localhost:8000/ping", nil)
	w := httptest.NewRecorder()

	ping(w, req)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, "pong", string(body[:]))
}