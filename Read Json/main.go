package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type File struct {
	Database struct {
		Host string `json: "host"`
		Port int    `json:"port"`
	} `json:"database`
	Host string `json: "host"`
	Port int    `json:"port"`
}

func LoadFile(fileName string) (File, error) {
	var config File
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func main() {
	fmt.Println("Starting the Aplication")
	file, _ := LoadFile("file.json")
	fmt.Println(file)
}
