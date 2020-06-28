package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////
var wg sync.WaitGroup
var c *kafka.Consumer

////////////////////////////////////////////////////////////////////
/*
	初始化工作
		初始化kafka环境
		初始化restful环境
		初始化定时线程，检查IPC状态
*/
func init() {
	go initKafka()
	go initRest()
	go timerTask()
	fmt.Println("init exit")
}

type kafkaconf struct {
	SERVERS string `yaml:"bootstrap.servers"`
	GROUP   string `yaml:"group.id"`
	OFFSET  string `yaml:"auto.offset.reset"`
}

func (c *kafkaconf) getKafkaConf() *kafkaconf {
	yamlFile, err := ioutil.ReadFile("kafka_conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func initKafka() {
	wg.Add(1)
	defer wg.Done()

	fmt.Println("init Kafka")

	var c kafkaconf
	conf := c.getKafkaConf()
	fmt.Println(conf.SERVERS)
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{

		"bootstrap.servers": conf.SERVERS,
		"group.id":          conf.GROUP,
		"auto.offset.reset": conf.OFFSET,
	})

	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{"SELECT"}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			var info ffprobeinfo
			info.mode = ONETIME
			info.timeout = 10
			info.streams = string(msg.Value)
			go schedFFprobe(&info)

		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	fmt.Println("out kafka")
}

func initRest() {
	wg.Add(1)
	defer wg.Done()

	fmt.Println("init Rest")

	r := gin.Default()

	r.GET("/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": `
			/c/status		采集器运行状态
			/c/close		采集器停止
			/d/check		检查IPC状态
			/d/select       启动IPC采集
			`,
		})
	})

	c := r.Group("/c")
	{
		c.GET("/status", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "status",
			})
		})

		c.GET("/close", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "close",
			})

			for {
				wg.Done()
			}
		})
	}

	d := r.Group("/d")
	{
		d.POST("/check", func(c *gin.Context) {
			var info ffprobeinfo
			info.mode = ONETIME
			info.timeout = 10
			info.streams = "rtmp://202.69.69.180:443/webcast/bshdlive-pc"
			status := schedFFprobe(&info)
			c.JSON(200, gin.H{
				"stream": info.streams,
				"status": status,
			})
		})

		d.POST("/select", func(c *gin.Context) {
			var info ffprobeinfo
			info.timeout = 10
			info.mode = ONETIME
			info.streams = "rtmp://202.69.69.180:443/webcast/bshdlive-pc"
			if status := schedFFprobe(&info); status == "living" {
				fmt.Println("ffmpeg " + info.streams)
				c.JSON(200, gin.H{
					"stream": info.streams,
					"select": "success",
				})

			} else {
				c.JSON(200, gin.H{
					"stream": info.streams,
					"select": "failed",
				})
			}
		})
	}

	r.Run(":8080")
}

func timerTask() {
	wg.Add(1)
	defer wg.Done()
	fmt.Println("timerTask")

	var info ffprobeinfo
	info.mode = PERIOD
	info.timeout = 10
	info.streams = "rtmp://202.69.69.180:443/webcast/bshdlive-pc"
	schedFFprobe(&info)
}

////////////////////////////////////////////////////////////////////
type ffmpegconf struct {
	TIMEOUT time.Duration `yaml:"ffmpeg.timeout"`
	TIMES   int           `yaml:"ffmpeg.times"`
}

func (c *ffmpegconf) getFFmpegConf() *ffmpegconf {
	yamlFile, err := ioutil.ReadFile("ffmpeg_conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func getFFmpegArgsAndCameras() (*os.File, *os.File) {
	conf1, err := os.Open("./argv.conf")
	defer conf1.Close()
	conf2, err := os.Open("./camera.conf")
	defer conf2.Close()

	if err != nil {
		return nil, nil
	}

	if err != nil {
		return nil, nil
	}

	return conf1, conf2
}

func schedFFmpeg() bool {
	wg.Add(1)
	defer wg.Done()

	var c1 ffmpegconf
	ffmpegconf := c1.getFFmpegConf()
	//conf1, conf2 := getFFmpegArgsAndCameras()
	conf1, _ := os.Open("./argv.conf")
	defer conf1.Close()
	conf2, _ := os.Open("./camera.conf")
	defer conf2.Close()
	a := bufio.NewScanner(conf1)
	c := bufio.NewScanner(conf2)

	a.Scan()
	for c.Scan() {
		fmt.Println(c.Text())
		func(a *bufio.Scanner, c *bufio.Scanner) {
			wg.Add(1)
			idAddr := strings.Split(c.Text(), " ")
			n := 1
			for n < ffmpegconf.TIMES {
				ffmpegArgs := fmt.Sprintf(a.Text(), idAddr[1])
				out, err := command(ffmpegconf.TIMEOUT, "ffmpeg", strings.Split(ffmpegArgs, " ")[:]...)
				if err != nil {
					fmt.Println(string(out[:]))
					return
				}
				n++
			}

			wg.Done()
		}(a, c)
	}

	return true
}

type Mode int

const (
	ONETIME Mode = iota
	PERIOD
)

type ffprobeinfo struct {
	mode    Mode
	timeout time.Duration
	streams string
}

func schedFFprobe(info *ffprobeinfo) (status string) {
	if info.mode == PERIOD {
		c := time.Tick(5 * time.Second)
		for {
			<-c
			go func() {
				out, _ := command(info.timeout, "ffprobe", strings.Split(info.streams, " ")[:]...)
				fmt.Println(string(out[:]))
				if strings.Contains(string(out[:]), "Stream") {
					fmt.Println(info.streams + " is living")
				} else {
					fmt.Println(info.streams + " is offline")
				}
			}()
		}
	} else {
		out, _ := command(info.timeout, "ffprobe", strings.Split(info.streams, " ")[:]...)
		fmt.Println(string(out[:]))
		if strings.Contains(string(out[:]), "Stream") {
			fmt.Println(info.streams + " is living")
			return "living"
		} else {
			fmt.Println(info.streams + " is offline")
			return "offline"
		}
	}

	return "offline"
}

func command(timeout time.Duration, name string, arg ...string) ([]byte, error) {
	ctxt, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	fmt.Println(arg)
	cmd := exec.CommandContext(ctxt, name, arg...)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return buf.Bytes(), err
	}

	if err := cmd.Wait(); err != nil {
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}

////////////////////////////////////////////////////////////////////
func main() {
	fmt.Println("ingest run!")
	schedFFmpeg()
	wg.Wait()
	fmt.Println("ingest closed!")
}
