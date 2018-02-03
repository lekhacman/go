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

func TestWriteText(t *testing.T) {
	path := "./write.txt"
	text := "Hello Man"
	path, err := WriteText(path, text)
	if err != nil {
		t.Errorf("Failed to write file")
	}
	assert.Equal(t, text, ReadContent(path))
}