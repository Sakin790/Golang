package main

import (
  "net/http"
  "fmt"

  "github.com/gin-gonic/gin"
)



func getting(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{
    "method": "Getting",
  })
}


func main() {
  // Create a Gin router with default middleware (logger and recovery)
  router := gin.Default()

  router.GET("/get", getting);
  router.GET("/ping", func(c *gin.Context) {
    fmt.Println(c)
    c.JSON(http.StatusOK, gin.H{
      "message": "Go server is healthy",
    })
  });

  router.GET("/", func(c *gin.Context){

    c.JSON(http.StatusOK, gin.H{
      "Message":"Home Route is OK",
    })
  })

  // Start server on port 8080 (default)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  router.Run()
}
