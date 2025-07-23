package main

import (
	"fmt"
	"os"
	"encoding/json"
)
func getConfigString(name string)(string){
  // this function looks for config.json and retrieves the value
  	file, openErr := os.Open("config.json")
	if openErr != nil {
		fmt.Println("Error opening config! Exiting...\n", openErr)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config map[string]string
  decodeErr := decoder.Decode(&config)
  if decodeErr != nil {
		fmt.Println("error decoding config! exiting...\n", decodeErr)
    os.Exit(1)
  }
  configString, configErr := config[name]
  if !configErr {
    fmt.Println("value does not exist in config.json! exiting...\n")
    os.Exit(1)
  }
  return configString


}

func getConfigInteger(name string)(int){
  // this function looks for config.json and retrieves the value

}
