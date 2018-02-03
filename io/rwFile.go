package io

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadContent(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)

	return str

}

func WriteText(filePath string, content string) (string, error) {
	var err error

	reader := strings.NewReader(content)

	b, err := ioutil.ReadAll(reader)

	if err != nil {
		return filePath, err
	}

	// perm param is a linux's file permission setting
	err = ioutil.WriteFile(filePath, b, 0644)

	return filePath, err
}