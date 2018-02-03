package io

import (
	"io/ioutil"
	"log"
)

func ReadContent(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)

	return str

}