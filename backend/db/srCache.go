package db

import (
	"fmt"
	"noAddit/utils"
)

func srRedis(){
  	cacheTimeout := utils.GetConfigInteger("cacheTimeoutHours")
	fmt.Print(cacheTimeout)
}