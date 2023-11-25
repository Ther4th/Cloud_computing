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

  hours, _, _ := time.Now().Clock()
  if(hours > 12 && hours < 18){
    c.IndentedJSON(http.StatusOK, "It is currently afternoon")
  }else{
    c.IndentedJSON(http.StatusOK, "It is currently not afternoon")

  }
}

func main() {
	port := "8082"

  router := gin.Default()
  router.GET("/isitafternoon", defaultFunc)
  address := ":" + port
  router.Run(address)

}
