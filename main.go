package main

import (
	_ "okgopro/internal/packed"

	_ "okgopro/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"okgopro/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
