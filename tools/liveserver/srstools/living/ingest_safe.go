package main

import (
	"bufio"
	"bytes"
	_ "bytes"
	"fmt"
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

func CheckHW() string {
	list := "cat /sys/class/dmi/id/product_uuid /sys/class/dmi/id/board_serial | tr -d '\n/-'"
	cmd := exec.Command("bash", "-c", list)
	w := bytes.NewBuffer(nil)
	cmd.Stdout = os.Stdout
	cmd.Stderr = w
	err := cmd.Run()
	if err != nil {
		//fmt.Printf("%s\n", w.Bytes())
		fmt.Printf("------------")
		return string(w.Bytes())
	} else {
		return ""
	}
}

func CheckEnv() bool {
	c, err := exec.Command("/bin/sh", "-c", "cat /proc/1/status", "|", "head -1").CombinedOutput()
	if err != nil {
		if strings.Contains(string(c), "systemd") {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
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
		return
	} else {
		fmt.Println(hwinfo)
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
