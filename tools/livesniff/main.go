package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	_ "github.com/go-echarts/go-echarts/charts"
)

var (
	h bool
	i string
	p string
	t int64
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
	flag.StringVar(&i, "i", "camera.conf", "流地址文件或流地址")
	flag.StringVar(&p, "p", "none", "流地址变化策略")
	flag.Int64Var(&t, "t", 10, "默认超时：10s")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}

	if strings.Index(i, "://") != -1 {

	} else {

	}

	//streamidRange := strings.Split(i, "/")
	//id := streamidRange[len(streamidRange)-1]
	if "none" == p {
		out, _ := Command("ffprobe", strings.Split("-i "+i, " ")[:]...)
		if strings.Index(out, "title") != -1 && strings.Index(out, "Stream") != -1 {
			fmt.Println("=======================================> stream is living")
		} else {
			fmt.Println("=======================================> stream is death")
		}
	}
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
