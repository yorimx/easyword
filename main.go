package main

import (
	_ "easyword/internal/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"easyword/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	g.I18n().SetLanguage("zh-CN")
	cmd.Main.Run(gctx.GetInitCtx())
}
