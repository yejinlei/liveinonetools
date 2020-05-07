package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	h  bool
	i  string
	ip string
	t  int64
	m  string
	l  int64

	lock sync.Mutex
	id   = make(chan string, 1000)
)

func usage() {
	fmt.Fprintf(os.Stderr, `livesniff version: livesniff/0.1
Usage: livesniff [-i]

Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&m, "m", "sniff", "sniff嗅探模式|stress压力测试模式|show显示模式")
	flag.StringVar(&i, "i", "camera.conf", "流地址文件或流地址")
	flag.Int64Var(&t, "t", 10, "默认超时：10s")

	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}

	/*动态生成、文件读取*/
	var liveURL []string
	if strings.Index(i, "://") != -1 {
		if strings.Contains(i, "^") {
			url := strings.Split(i, "^")
			urlRange := strings.Split(url[1], "-")
			s, _ := strconv.Atoi(urlRange[0])
			e, _ := strconv.Atoi(urlRange[1])
			for s <= e {
				liveURL = append(liveURL, url[0]+strconv.Itoa(s)+url[2])
				s++
			}
		} else {
			liveURL = append(liveURL, i)
		}
	} else {

	}

	/*选择模式，嗅探、压力、显示*/
	for _, v := range strings.Split(m, ";") {
		if "sniff" == v {
			for _, u := range liveURL {
				go func(u string) {
					out, _ := command("ffprobe", strings.Split("-i "+u, " ")[:]...)
					if strings.Index(out, "Stream") != -1 {
						id <- u + ":living"
					} else {
						id <- u + ":death"
					}
				}(u)
			}
		} else if "stress" == v {
			for _, u := range liveURL {
				go func(u string) {
					out, _ := command("ffmpeg", strings.Split("-c:v copy -c:a copy -f flv 1.flv -i "+u, " ")[:]...)
					if strings.Index(out, "Stream") != -1 {
						id <- u + ":living"
					} else {
						id <- u + ":death"
					}
				}(u)
			}
		} else if "show" == v {
			//for _, u := range liveURL {
			//	go pipeCommand(u)
			//}
			for _, u := range liveURL {
				go func(u string) {
					command("ffplay", strings.Split("-i "+u, " ")[:]...)
				}(u)
			}
		}
	}

	stat()
	//time.Sleep(100000000000)
}

func command(name string, arg ...string) (string, error) {
	Timeout := time.Duration(t) * time.Second
	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctxt, name, arg...)
	//cmd.Stdout = os.Stdout
	if out, err := cmd.CombinedOutput(); err != nil {
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			return "", errors.New("command timeout")

		}
		return string(out), err
	} else {
		return string(out), nil
	}
}

func pipeCommand(u string) {
	var commands []*exec.Cmd
	sysType := runtime.GOOS
	shellType := map[string][]string{"linux": []string{"sh", "-c"},
		"windows": []string{"cmd", "/C"}}
	shellArgs := shellType[sysType]
	//ffmpegArgs := "ffmpeg -i rtmp://10.10.82.66:11936/live/5021 -i rtmp://10.10.82.66:11936/live/6058 -filter_complex 'nullsrc=size=1024x768[base]; [0:v]setpts=PTS-STARTPTS,scale=517x384[l1]; [1:v]setpts=PTS-STARTPTS,scale=517x384[l2]; [base][l1]overlay=x=0:y=0[tmp1]; [tmp1][l2]overlay=x=0:y=384' -f flv -"
	ffmpegArgs := "ffmpeg -t 10 -i " + u + " -vcodec copy -f flv -"
	ffplayArgs := "ffplay -i -"
	commands = append(commands, exec.Command(shellArgs[0], shellArgs[1], ffmpegArgs))
	commands = append(commands, exec.Command(shellArgs[0], shellArgs[1], ffplayArgs))
	for i := 1; i < len(commands); i++ {
		commands[i].Stdin, _ = commands[i-1].StdoutPipe()
	}
	commands[len(commands)-1].Stdout = os.Stdout
	for i := 1; i < len(commands); i++ {
		err := commands[i].Start()
		if err != nil {
			panic(err)
		}
	}

	commands[0].Run()

	for i := 1; i < len(commands); i++ {
		err := commands[i].Wait()
		if err != nil {
			panic(err)
		}
	}
}

func stat() {
	for {
		select {
		case urlID := <-id:
			fmt.Println(urlID)
		}
	}
}
