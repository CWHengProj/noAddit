package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)
func CallGETEndpoint(baseUrl string, subreddit string, limit string, timeFrame string)(json.RawMessage){
	fmt.Print("calling get endpoint...\n")
	url := baseUrl+subreddit+".json"+"?limit="+limit+"&t="+timeFrame
	fmt.Println(url)
	response, httpErr := http.Get(url)
	if httpErr != nil {
		fmt.Println("Error retrieving subreddit. Check if it exists. Exiting....", httpErr)
		os.Exit(1)
	}
	defer response.Body.Close()
	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		fmt.Println("Error reading response body. Exiting...",readErr)
		os.Exit(1)
	}
	return json.RawMessage(body)
}