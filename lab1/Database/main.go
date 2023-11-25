package main

import (
	//"time"
	"fmt"
  "database/sql"
	"log"
	"net/http"
  "github.com/gin-gonic/gin"
  _ "github.com/mattn/go-sqlite3"
)

type Searches struct {
    id     int
    firstname string
    surname string
}

func getDatabase(c *gin.Context){
  param := c.Param("name")
  db, err := sql.Open("sqlite3", "/app/data/database.db")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  results := ""
    row, err := db.Query("SELECT * FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer row.Close()
    for row.Next() { // Iterate and fetch the records from result cursor
        item := Searches{}
        err := row.Scan(&item.id, &item.firstname, &item.surname)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Print("gets here!\n")
        if(param == item.firstname){
          results += string(item.id) + " " + item.firstname + " " + item.surname + "\n"
        }
        //searches = append(searches, item)
    }



  db.Close()

  c.JSON(http.StatusOK, results)
}


func putDatabase(c *gin.Context){
  db, err := sql.Open("sqlite3", "/app/data/database.db")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()
  _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT, surname TEXT)")
  if err != nil {
    log.Fatal(err)
  }

  firstname := c.Param("firstname")
  surname := c.Param("surname")

  _, err = db.Exec("INSERT INTO users (firstname, surname) VALUES (?, ?)", firstname, surname)
	if err != nil {
		log.Fatal(err)
	}



  db.Close()
  c.JSON(http.StatusOK, "Insertion successful!")
}

func main() {
	port := "8083"

  db, err := sql.Open("sqlite3", "/app/data/database.db")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT, surname TEXT)")
  if err != nil {
    log.Fatal(err)
  }
  db.Close()


  router := gin.Default()
  rgroup := router.Group("/database")
  {
    rgroup.GET("/get/:name", getDatabase)
    rgroup.GET("/put/:firstname/:surname", putDatabase)

  }



  address := ":" + port
  router.Run(address)

}
