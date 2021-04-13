package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type PathDefinition struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func main() {
	jsonFile, err := os.Open("mail.json")
	if err != nil {
		writeDefaultFile()
		fmt.Println("\nThe mail.json file was invalid or not found. Please check the created json file and fill it with your paths.\n")
		return
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var definitions []PathDefinition

	err = json.Unmarshal(byteValue, &definitions)
	if err != nil {
		panic(err)
	}

	for i := range definitions {
		process(definitions[i].From, definitions[i].To)
	}
}

func writeDefaultFile() {
	file, _ := json.MarshalIndent([]PathDefinition{
		PathDefinition{
			From: "source/demo.html",
			To:   "out/demo.html",
		},
	}, "", " ")
	_ = ioutil.WriteFile("mail.json", file, 0644)
}

func process(from string, to string) {
	f, err := os.Open(from)

	if err != nil {
		fmt.Println("Failed to open " + from + ", skipping file")
		return
	}

	byteValue, _ := ioutil.ReadAll(f)

	response, err := http.PostForm("https://templates.mailchimp.com/services/inline-css/", url.Values{
		"html": {string(byteValue)}})

	if err != nil {
		fmt.Println("Failed to process " + from + ", skipping file")
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Failed to process " + from + ", skipping file")
		return
	}

	_ = os.MkdirAll(filepath.Dir(to), os.ModePerm)

	err = ioutil.WriteFile(to, body, 0644)
	if err != nil {
		fmt.Println("Failed to write " + from + ", skipping file")
		return
	}

	log.Println("Converted " + from)
}
