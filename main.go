package main

import (
	_ "okgopro/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"okgopro/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
