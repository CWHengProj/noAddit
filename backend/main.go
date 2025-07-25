package main

import (
  "noAddit/controllers"
	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

func main(){
  routes := gin.Default()
  //TODO: cors for my frontend once confirmed
  routes.Use(cors.Default())
  // healthcheck
  routes.GET("/ping", controllers.Ping)
  // subreddits
  routes.GET("/news", controllers.Subreddits)
  routes.Run()
    
  }
