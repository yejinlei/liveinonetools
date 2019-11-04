package main

import (
	"bytes"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	_ "github.com/tedcy/fdfs_client"
	_ "log"
	"os/exec"
	_ "strings"
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
	for {
		fmt.Println("kafka_service")
		fetch_msg()
		ffmpeg_warp()
	}
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

func ffmpeg_warp() {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	//list := strings.Split(ffmpeg_args, " ")
	for {
		ffmpeg_args := "ffmpeg -i rtsp://wowzaec2demo.streamlock.net/vod/mp4:BigBuckBunny_115k.mov -r 1 -f image2 image%d.jpg"
		cmd := exec.Command("bash", "-c", ffmpeg_args)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			continue
		}

		fmt.Println(stdout)
		fmt.Println("finish")
	}
}

func fetch_msg() {
	fmt.Println("fetch_msg")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return

	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("kafka_go_test", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return

	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())

		}

	}
}
