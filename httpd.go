package main

import (
  "os"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.String(200, "pong")
  })

  r.Run(":" + os.Getenv("PORT"))
}
