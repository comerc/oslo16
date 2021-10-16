package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	result := make(map[int64]string)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fileName := f.Name()
		a := strings.Split(fileName, ".")
		if len(a) != 2 || a[1] != "json" {
			continue
		}
		file, err := os.Open("./files/" + fileName)
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
		result[fileData.AppId] = fileData.AppHash
	}
	fmt.Printf("%v", result)
	fmt.Print("len: ", len(result))
}
