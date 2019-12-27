package main

import (
	"bufio"
	_ "bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/pkg/profile"
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

func main() {

	var wg sync.WaitGroup
	defer profile.Start(profile.MemProfile).Stop()
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
