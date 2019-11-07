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

	_ "github.com/go-echarts/go-echarts/charts"
)

var (
	h  bool
	i  string
	ip string
	t  int64
	m  string

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
	var liveUrl []string
	if strings.Index(i, "://") != -1 {
		url := strings.Split(i, "/")
		urlLen := len(url)
		if strings.Contains(url[urlLen-1], "-") {
			urlRange := strings.Split(url[urlLen-1], "-")
			s, _ := strconv.Atoi(urlRange[0])
			e, _ := strconv.Atoi(urlRange[1])
			for s <= e {
				liveUrl = append(liveUrl, strings.Join(url[:urlLen-1], "/")+"/"+strconv.Itoa(s))
				s += 1
			}
		} else {
			liveUrl = append(liveUrl, i)
		}
	} else {

	}

	/*根据策略生成地址池*/
	//streamidRange := strings.Split(i, "/")
	//id := streamidRange[len(streamidRange)-1]

	//liveUrl = []string{"rtmp://10.10.82.66:11936/live/6058", "rtmp://10.10.82.66:11936/live/5021"}

	/*选择模式，嗅探、压力、显示*/
	for _, v := range strings.Split(m, ";") {
		if "sniff" == v {
			for _, u := range liveUrl {
				go func(u string) {
					out, _ := Command("ffprobe", strings.Split("-i "+u, " ")[:]...)
					if strings.Index(out, "title") != -1 && strings.Index(out, "Stream") != -1 {
						id <- u + ":living"
					} else {
						id <- u + ":death"
					}
				}(u)
			}
		} else if "stress" == v {

		} else if "show" == v {
			go PipeCommand()
		}
	}

	stat()
	//time.Sleep(100000000000)
}

func Command(name string, arg ...string) (string, error) {
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

func PipeCommand() {
	var commands []*exec.Cmd
	sysType := runtime.GOOS
	shellType := map[string][]string{"linux": []string{"sh", "-c"},
		"windows": []string{"cmd", "/C"}}
	shell_args := shellType[sysType]
	ffmpeg_args := "ffmpeg -i rtmp://10.10.82.66:11936/live/5021 -i rtmp://10.10.82.66:11936/live/6058 -filter_complex 'nullsrc=size=1024x768[base]; [0:v]setpts=PTS-STARTPTS,scale=517x384[l1]; [1:v]setpts=PTS-STARTPTS,scale=517x384[l2]; [base][l1]overlay=x=0:y=0[tmp1]; [tmp1][l2]overlay=x=0:y=384' -f flv -"
	ffplay_args := "ffplay -i -"
	commands = append(commands, exec.Command(shell_args[0], shell_args[1], ffmpeg_args))
	commands = append(commands, exec.Command(shell_args[0], shell_args[1], ffplay_args))

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
		case urlId := <-id:
			fmt.Println(urlId)
		}
	}
}
