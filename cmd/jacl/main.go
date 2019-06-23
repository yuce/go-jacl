package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yuce/go-jacl"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: jacl configuration-file")
	}
	text, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	m := map[string]interface{}{}
	err = jacl.Unmarshal(string(text), &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

}
