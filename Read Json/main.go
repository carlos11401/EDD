package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type File struct {
	// List of dates
	Dates []struct {
		Index string `json:"Indice"`
		// list of Departments
		Department []struct {
			Name string `json:"Nombre"`
			// list of Stores
			Store []struct {
				Name          string `json:"Nombre"`
				Description   string `json:"Descripcion"`
				Contact       string `json:"Contacto"`
				Qualification int    `json:"Calificacion"`
			} `json:"Tiendas"`
		} `json:"Departamentos"`
	} `json:"Datos"`
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
	fmt.Println("Starting the Application")
	file, _ := LoadFile("file.json")
	fmt.Println(file)
}
