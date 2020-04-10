package main

import (
	"flag"
	"fmt"

	commonfunc "github.com/jin123/mocke-server/src/common"
	"github.com/jin123/mocke-server/src/vendor"
)

func main() {
	port := flag.Int("port", 10000, "this is http port")
	fmt.Println("当前端口是=", *port)
	commonfunc.GetCurrentPath()
	vendor.CreateHttpConnect(*port)

}
