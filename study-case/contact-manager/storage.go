package main

import (
	"os"
	"fmt"
	"encoding/json"
)

var path string = "./contacts.json"

func loadContacts() {
	checkPath()

	fileContents := openFile()

	var jsonContents map[string][]Contact

	if err := json.Unmarshal(fileContents, &jsonContents); err != nil {
		fmt.Println(err.Error())
		return
	}

	contacts = jsonContents["contacts"]
}

func checkPath() {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func openFile() []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return data
}
