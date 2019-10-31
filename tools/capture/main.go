package main

import (
	"fmt"
	_ "github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	_ "github.com/tedcy/fdfs_client"
	"sync"
)

var count = 1

func restful_service() {
	r := gin.Default()

	capture := r.Group("/capture/")
	{
		capture.GET("", capture_get)
		capture.POST("", capture_post)
	}

	r.Run(":28005")
	count -= 1
}

func kafka_service() {
	fmt.Println("kafka")
	count -= 1
}

func main() {
	var wg sync.WaitGroup
	go restful_service()
	go kafka_service()
	wg.Add(count)
	wg.Wait()
}

func capture_get(c *gin.Context) {
	fmt.Println("capture_get")
}

func capture_post(c *gin.Context) {
	fmt.Println("capture_post")
}
