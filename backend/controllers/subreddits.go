package controllers

import (
	"fmt"
	"net/http"
	"noAddit/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func Subreddits(c *gin.Context){
	fmt.Println("Retrieving config...")
	  // retrieve values from config
  	baseUrl := utils.GetConfigString("baseUrl")
  	timeFrame := utils.GetConfigString("timeFrame")
  	// cacheTimeout := utils.GetConfigInteger("cacheTimeoutHours")
  	maxPostsAllowed := utils.GetConfigInteger("maxPosts")

	// Input validation
	subreddits := c.Query("subreddits")
	if subreddits == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"Missing query for subreddits"})
		return 
	}
	posts := c.Query("limit")
	if posts == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"Missing query for limit"})
		return 
	}
	limit, postsErr := strconv.Atoi(posts)
	if postsErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"Limits should be an integer value"})
		return 
	}
	subredditArray := strings.Split(subreddits, "+")
	if len(subredditArray) > 5 {
		subredditArray = subredditArray[0:5]
	}
	if limit > maxPostsAllowed {
		limit = maxPostsAllowed
	}
	limitStr := strconv.Itoa(limit)
	var results []json.RawMessage
	for i := 0; i < len(subredditArray); i++ {
		srData := utils.CallGETEndpoint(baseUrl,subredditArray[i],limitStr,timeFrame)
		results = append(results, srData)
	}
	// TODO: call the service, retrieve the data (or cached data), return as an array

	// Expose the endpoint with crafted subreddit json information to display
	c.JSON(http.StatusOK, gin.H{"Subreddit information":results})
}