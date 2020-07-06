package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		data, err := ioutil.ReadFile("/data/body")
		c.JSON(200, gin.H{
			"message": "Hello from zeet, post anything to /write",
			"data":    string(data),
			"error":   err,
		})
	})
	r.POST("/write", func(c *gin.Context) {
		bytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		err = ioutil.WriteFile("/data/body", bytes, 0644)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
