package utils

import (
	"fmt"
	"os"
	"encoding/json"
	"strconv"
)
func GetConfigString(name string)(string){
  	file, openErr := os.Open("config/config.json")
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
    fmt.Println("value does not exist in config.json! exiting...\n", configErr)
    os.Exit(1)
  }
  return configString


}

func GetConfigInteger(name string)(int){
  configString := GetConfigString(name)
  num, numErr := strconv.Atoi(configString)
  if numErr != nil {
	fmt.Println("value is not an integer! Please check your config before trying again. Exiting...\n",numErr)
	os.Exit(1)
  }
  return num
}
