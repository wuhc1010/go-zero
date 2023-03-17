package main

import (
	"github.com/wuhc1010/go-zero/core/load"
	"github.com/wuhc1010/go-zero/core/logx"
	"github.com/wuhc1010/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
