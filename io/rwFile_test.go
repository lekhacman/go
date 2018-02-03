package io

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestReadContent(t *testing.T) {
	content := ReadContent("./text.txt")
	fmt.Println(content)
	assert.Equal(t, "Hello World!\n", content)
}
