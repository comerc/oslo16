package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type FileData struct {
	AppId   int64  `json:"app_id"`
	AppHash string `json:"app_hash"`
}

func main() {
	files, err := ioutil.ReadDir("./files")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		file, err := os.Open("./files/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		fileData := FileData{}
		if err := json.Unmarshal(data, &fileData); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", fileData)
	}
}
