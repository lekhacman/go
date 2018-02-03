package concurrency

import (
	"testing"
	"fmt"
)

func TestGetPingPong(t *testing.T) {
	getResult := GetPingPong()
	fmt.Println(getResult)
	if getResult != "pingpong" && getResult != "pongping" {
		t.Errorf("Incorrect response: %s", getResult)
	}
}
