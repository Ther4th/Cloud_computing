package main

import (
  "io/ioutil"
	"fmt"
	"log"
	"net/http"
	"os"
  "github.com/gin-gonic/gin"
)

func getDatabase(c *gin.Context){
  param := c.Param("name")
  url := "http://localhost:8083/database/get/" + param
  resp, err := http.Get(url)
  if err != nil {
   log.Fatalln(err)
 }
 body, err := ioutil.ReadAll(resp.Body)
 if err != nil {
   log.Fatalln(err)
 }
 sb := string(body)
  if (resp.StatusCode == 200){

    c.IndentedJSON(http.StatusOK, sb)
  }else{
    c.IndentedJSON(http.StatusOK, "Failed to connect to serivce!")
 }

}

func putDatabase(c *gin.Context){
  firstname := c.Param("firstname")
  surname := c.Param("surname")
  url := "http://localhost:8083/database/put/" + firstname + "/" + surname

  resp, err := http.Get(url)
  if err != nil {
   log.Fatalln(err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
   log.Fatalln(err)
  }
  sb := string(body)
  if (resp.StatusCode == 200){
    c.IndentedJSON(http.StatusOK, sb)
  }else{
    c.IndentedJSON(http.StatusOK, "Failed to connect to serivce!")
  }

}

func defaultFunc(c *gin.Context){
  fmt.Print("User connected to home page\n")
  c.IndentedJSON(http.StatusOK, "Welcome to my cloud solution")
}

func timeFunc(c *gin.Context){
  fmt.Print("Time was called\n")
  resp, err := http.Get("http://localhost:8081/time")
  if err != nil {
   log.Fatalln(err)
 }

  message := "Failed to connect"
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  sb := string(body)
   if (resp.StatusCode == 200){
     message = "current date and time is: " + sb
     c.IndentedJSON(http.StatusOK, message)
   }else{
     c.IndentedJSON(http.StatusOK, message)
  }
}

func afternoonFunc(c *gin.Context){
  fmt.Print("Time was called\n")
  resp, err := http.Get("http://localhost:8082/isitafternoon")
  if err != nil {
   log.Fatalln(err)
 }

  message := "Failed to connect"
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  sb := string(body)
   if (resp.StatusCode == 200){

     c.IndentedJSON(http.StatusOK, sb)
   }else{
     c.IndentedJSON(http.StatusOK, message)
  }
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

  router := gin.Default()
  router.GET("/", defaultFunc)

  router.GET("/time", timeFunc)
  router.GET("/afternoon", afternoonFunc)
  rgroup := router.Group("/database")
  {
    rgroup.GET("/get/:name", getDatabase)
    rgroup.GET("/put/:firstname/:surname", putDatabase)
  }

  address := ":" + port
  router.Run(address)
}
