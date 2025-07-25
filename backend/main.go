package main

import (
  "noAddit/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
  routes := gin.Default()
  // healthcheck
  routes.GET("/ping", controllers.Ping)
  // subreddits
  routes.GET("/news", controllers.Subreddits)
  routes.Run()
    
  }
