package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	_ "strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	_ "github.com/tedcy/fdfs_client"
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
		go consumer_msg()
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
	for {
		select {
		case <-CS:
			ffmpeg_args := "-i rtsp://wowzaec2demo.streamlock.net/vod/mp4:BigBuckBunny_115k.mov -r 1 -f image2 image%d.jpg"
			arg := strings.Split(ffmpeg_args, " ")
			out, err := Command("ffmpeg", arg[:]...)
			if err != nil {
				log.Println("错误返回:", err.Error())
			} else {
				log.Println("正常返回:", out)
			}

			fmt.Println(out)
			fmt.Println("finish")
		}
	}
}

func init() {
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
}

var CS = make(chan bool, 1)

func consumer_msg() {
	fmt.Println("consumer_msg")

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 1
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Metadata.Full = true
	config.Version = sarama.V0_10_0_0
	config.ClientID = "sasl_scram_client"
	config.Metadata.Full = true

	// consumer
	consumer, err := sarama.NewConsumer([]string{"10.10.82.66:9092"}, config)
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

	c := time.Tick(20 * time.Second)
	for {
		select {
		//case msg := <-partition_consumer.Messages():
		//	fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
		//		msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		//case err := <-partition_consumer.Errors():
		//	fmt.Printf("err :%s\n", err.Error())

		case <-c:
			CS <- true
			fmt.Println("cs <- true")
		}

	}
}

var Timeout = 10 * time.Second

func Command(name string, arg ...string) (string, error) {
	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctxt, name, arg...)
	if out, err := cmd.Output(); err != nil {
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			return "", errors.New("command timeout")

		}
		return string(out), err
	} else {
		return string(out), nil
	}
}
