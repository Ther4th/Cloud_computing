package main

import (
	"time"
	"fmt"
	//"log"
	"net/http"
  "github.com/gin-gonic/gin"
)


func defaultFunc(c *gin.Context){
  fmt.Print("returns 200!")
	now := time.Now()
  c.IndentedJSON(http.StatusOK, now.Format("15:04:05"))
}

func main() {
	port := "8081"

  router := gin.Default()
  router.GET("/time", defaultFunc)
  address := ":" + port
  router.Run(address)

}
