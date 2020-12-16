package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func status(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"status":"ok", "code":http.StatusOK})
}

func main() {
    port := ":4000"

    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    router.LoadHTMLGlob("templates/*")
    router.Static("/assets", "./assets")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title":"Gin Runner",
            "environment":"development",
        })
    })

    router.GET("/status", status)

    // POST("/shutdown/:token", shutdown)

    router.Run( port )
}
