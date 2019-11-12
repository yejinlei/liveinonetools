package main

import (
	"fmt"

	_ "github.com/gwuhaolin/livego/configure"
	_ "github.com/gwuhaolin/livego/protocol/hls"
	_ "github.com/gwuhaolin/livego/protocol/httpflv"
	_ "github.com/gwuhaolin/livego/protocol/httpopera"
	_ "github.com/gwuhaolin/livego/protocol/rtmp"
)

func main() {
	fmt.Println("livego wasm")
}
