package main

import (
	"bufio"
	_ "bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func warp(ffmpeg_args string) {
	list := strings.Split(ffmpeg_args, " ")
	for {
		cmd := exec.Command(list[0], list[1:]...)
		//w := bytes.NewBuffer(nil)
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = w
		cmd.Run()
		//if err != nil {
		//      fmt.Printf("%s\n", w.Bytes())
		//}
		time.Sleep(time.Second * 5)
	}
}

//cat /sys/class/dmi/id/product_uuid | tr '[:upper:]' '[:lower:]' | openssl md5
func CheckHW() string {
	c, err := exec.Command("cat", "/sys/class/dmi/id/product_uuid").CombinedOutput()
	if err != nil {
		return ""
	}

	hash := md5.New()
	io.WriteString(hash, strings.ToLower(string(c)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func CheckEnv() bool {
	c, err := exec.Command("cat", "/proc/1/status").CombinedOutput()
	if err != nil {
		return false
	}

	if strings.Contains(string(c), "systemd") {
		return false
	} else {
		return true
	}
}

func CheckRegcode(regcode string) bool {
	var regCodeFile string
	regCodeFile = os.Getenv("REGCODEFILE")
	r, _ := os.Open(regCodeFile)
	defer r.Close()
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if regcode == line {
			return true
		}
	}

	return false
}

func main() {

	isInDocker := CheckEnv()
	if isInDocker == false {
		fmt.Println(`   本计算机程序受版权法和国际条约保护，如未经授权而擅自复制、传播、
修改本程序（或者其中任何部分）将受到严厉的刑事及民事制裁，并在法律
允许的范围内受到最大可能的起诉。`)
		return
	}

	hwinfo := CheckHW()
	if hwinfo == "" {
		fmt.Println("无法获取本机信息")
		return
	}

	if CheckRegcode(hwinfo) == false {
		fmt.Println("非法注册码")
		return
	}

	var wg sync.WaitGroup
	//defer profile.Start(profile.MemProfile).Stop()
	conf1, err := os.Open("./argv.conf")
	if err != nil {
		defer conf1.Close()
		return
	}

	conf2, err := os.Open("./camera.conf")
	if err != nil {
		defer conf2.Close()
		return
	}

	a := bufio.NewScanner(conf1)
	a.Scan()

	c := bufio.NewScanner(conf2)
	count := 1
	for c.Scan() {
		id_addr := strings.Split(c.Text(), " ")
		//fmt.Println(fmt.Sprintf(a.Text(), id_addr[1], id_addr[0]))
		ffmpeg_args := "ffmpeg" + " " + fmt.Sprintf(a.Text(), id_addr[1], id_addr[0])
		go warp(ffmpeg_args)
		count += 1
	}

	wg.Add(count)
	wg.Wait()
}
