package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	_ "strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/tedcy/fdfs_client"
	"gopkg.in/ini.v1"
)

var count int = 1

var http_request_count = 0
var task_count = 0

var locker_for_gor sync.Mutex

var request_body = make(chan string, 150)

//var ini_full_path = "D:\\my_work\\yejinlei_code\\work_for_git\\yejinlei\\liveinonetools\\tools\\capture\\my.ini"
var ini_full_path = "my.ini"

var producer sarama.SyncProducer
var consumer sarama.Consumer
var kafka_addr_for_producer string

var kafka_recv_topic_name = "TOPIC_NB_VIDEO_RECON"
var kafka_feedback_topic_name = "TOPIC_NVR_DOOR_RECON"

var rtmp_server_ipport string

var ffmpeg_full_path string
var image_perfix_path string

var fdfs_client_obj *fdfs_client.Client
var fdfs_conf_file_path string
var fdfs_conf_server_addr string

func init() {
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)

	//conf
	cfg, err := ini.Load(ini_full_path)
	if err != nil {
		log.Println("ini open fail err:", err.Error())
		os.Exit(1)
	}

	rtmp_server_ipport = cfg.Section("common").Key("rtmp_server_addr").String()

	image_perfix_path = cfg.Section("common").Key("image_save_path").String()
	if runtime.GOOS == "windows" {
		image_perfix_path += "\\"
	} else {
		image_perfix_path += "/"
	}

	fdfs_conf_file_path = cfg.Section("fdfs").Key("conf_full_path").String()
	fdfs_conf_server_addr = cfg.Section("fdfs").Key("server_addr").String()

	kafka_addr_for_producer = cfg.Section("kafka").Key("server_addr").String()

	ffmpeg_full_path = cfg.Section("ffmpeg").Key("bin_full_path").String()

	//
	consumer, err = sarama.NewConsumer([]string{kafka_addr_for_producer}, nil)
	if err != nil {
		log.Fatalln(err)
		consumer = nil
	}

	//var err error = nil
	producer, err = sarama.NewSyncProducer([]string{kafka_addr_for_producer}, nil)
	if err != nil {
		log.Fatalln(err)
		producer = nil
	}

	fdfs_client_obj, err = fdfs_client.NewClientWithConfig(fdfs_conf_file_path)
	if err != nil {
		log.Fatalln(err)
		fdfs_client_obj = nil
	}
}

func clean_work() {
	if consumer != nil {
		consumer.Close()
	}

	if producer != nil {
		producer.Close()
	}

	if fdfs_client_obj != nil {
		fdfs_client_obj.Destory()
	}
}

func main() {
	var wg sync.WaitGroup

	//go restful_service()
	go image_kafka_service(kafka_recv_topic_name)

	defer clean_work()

	wg.Add(count)
	wg.Wait()
}

func restful_service() {
	r := gin.Default()

	capture := r.Group("/capture")
	{
		capture.GET("", capture_get)
		capture.POST("", capture_post)
	}

	//go image_save_worker()

	r.Run(":28005")
	count -= 1
}

func image_kafka_service(topic_name string) {
	//topic
	partitionConsumer, err := consumer.ConsumePartition(topic_name, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln(err)
		return
	}

	/*consumed := 0
	ConsumerLoop:*/
	for {
		select {
		case msg := <-partitionConsumer.Messages():

			msg_body := string(msg.Value)
			if len(msg_body) > 0 {
				log.Printf("Consumed message offset %d, msg: %s\n", msg.Offset, string(msg.Value))
				go image_save_worker(msg_body, topic_name)
			}

			//log.Printf("Consumed message offset %d, msg: %s\n", msg.Offset, string(msg.Value))
			//consumed++
			/*case <-signals:
			break ConsumerLoop*/
		}
	}
}

//////////////////////////////////////////////////////////

func capture_get(c *gin.Context) {
	fmt.Println("capture_get")
}

func capture_post(c *gin.Context) {
	if http_request_count > 1000*1000 {
		http_request_count = 0
	}

	http_request_count++

	fmt.Println("capture_post http_request_count:", http_request_count, "cur chan len: ", len(request_body))

	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	body_string := string(x)

	//fmt.Println(body_string)

	c.PureJSON(200, gin.H{
		"Hello": "world!",
	})

	c.Writer.Flush()

	go image_save_worker(body_string, kafka_recv_topic_name)

	//request_body <- body_string
	//nofuckblock_chanwrite(request_body, body_string)
}

//////////////////////////////////////////////////////

type image_save_msg struct {
	messgeId  string
	equipCode string
}

func image_save_worker(body_string, recv_topic_name string) {
	/*for*/ {
		//body_string := <-request_body

		locker_for_gor.Lock()
		if task_count > 1000*1000 {
			task_count = 0
		}

		task_count++
		cur_task_id := strconv.Itoa(task_count)
		cur_task_id = "[" + cur_task_id + "]"
		locker_for_gor.Unlock()

		log.Println(cur_task_id, " image save start")

		/*
			{
				messgeId: '',
					equipCode: '',
					instruct: '',
				apiSource: 'nbVideo'
			}

		*/
		json_kv := make(map[string]interface{})
		json.Unmarshal([]byte(body_string), &json_kv)

		var msg image_save_msg

		msg.messgeId = json_kv["MessgeId"].(string)
		msg.equipCode = json_kv["equipCode"].(string)

		log.Println(cur_task_id, " message: ", body_string)

		file_url, err := do_image_save(cur_task_id, msg)
		if err != nil {
			log.Println(cur_task_id, " image save have some err: ", err.Error())
			return
		}

		json_kv["imgUrl"] = file_url

		kafka_feedback_topic_name_tmp := kafka_feedback_topic_name
		if recv_topic_name == kafka_recv_topic_name {
			kafka_feedback_topic_name_tmp = kafka_feedback_topic_name
			json_kv["apISource"] = "NVR_DOOR"
		}

		newjson_str, err := json.Marshal(json_kv)

		image_kafka_msg_feedback(cur_task_id, string(newjson_str), kafka_feedback_topic_name_tmp)

		log.Println(cur_task_id, " image save end")
	}
}

func do_image_save(cur_task_id string, msg image_save_msg) (string, error) {
	cur_time := strconv.FormatInt(time.Now().Unix(), 10)

	file_name := "image_" + msg.equipCode + "_" + cur_time + "_" + cur_task_id + ".jpg"
	file_full_path := image_perfix_path + file_name

	var command_param string
	command_param = "-i rtmp://" + rtmp_server_ipport + "/live/" + msg.equipCode
	command_param += " -frames:v 1 -y -f image2 "
	command_param += file_full_path

	log.Println(cur_task_id, " run ffmpeg now")

	err, erroutput := exec_ffmpeg_command(cur_task_id, command_param)
	if err != nil {
		log.Println(cur_task_id, " ffmpeg run error:", err.Error(), "err string: ", erroutput)
		return "", errors.New("ffmpeg run error")
	}

	file_url, err := commit_file_to_fdfs(cur_task_id, file_full_path)
	if err != nil {
		log.Println(cur_task_id, " commit file error:", err.Error())
		return "", errors.New("commit file fail")
	}
	log.Println(cur_task_id, " file_url: ", file_url, ", file_full_path: ", file_full_path)

	os.Remove(file_full_path)

	return file_url, nil
}

func exec_ffmpeg_command(cur_task_id string, command_param string) (error, string) {
	//ffmpeg_args := "-i rtmp://10.10.82.66:11936/live/6002 -r 1 -f image2 image%d.jpg"
	//arg := strings.Split(command_param, " ")
	arg := strings.Fields(command_param)
	out, err, stderr_string := Command(ffmpeg_full_path, arg[:]...)
	/*if err != nil {
		log.Println("错误返回:", err.Error())
	} else {
		log.Println("正常返回:", out)
	}*/

	fmt.Println(out)
	log.Println(cur_task_id, " finish")
	return err, stderr_string
}

func commit_file_to_fdfs(cur_task_id string, file_full_path string) (string, error) {
	locker_for_gor.Lock()
	defer locker_for_gor.Unlock()

	if fdfs_client_obj == nil {
		return "", errors.New("err")
	}

	file_name_in_server, err := (*fdfs_client_obj).UploadByFilename(file_full_path)
	if err != nil {
		return "", err
	}

	file_url := "http://" + fdfs_conf_server_addr + "/" + file_name_in_server

	return file_url, nil
}

func image_kafka_msg_feedback(cur_task_id, kafka_msg, kafka_topic string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = kafka_topic
	msg.Key = sarama.StringEncoder("utf-8")
	msg.Partition = 0
	msg.Value = sarama.StringEncoder(kafka_msg)

	locker_for_gor.Lock()
	defer locker_for_gor.Unlock()
	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println(cur_task_id, " have some err ", err.Error())
	}

	log.Println(cur_task_id, " pid ", pid, ", offset:", offset)
}

///////////////////////////////////////////////////////////////////////////////////

var Timeout = 30 * time.Second

func Command(name string, arg ...string) (string, error, string) {
	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctxt, name, arg...)
	var errbuf bytes.Buffer
	cmd.Stderr = &errbuf
	if out, err := cmd.Output(); err != nil {
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			return "", errors.New("command timeout"), errbuf.String()

		}
		return string(out), err, errbuf.String()
	} else {
		return string(out), nil, ""
	}
}
