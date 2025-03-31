package iot

import (
	"embed"
	"github.com/busy-cloud/boat/menu"
	"github.com/busy-cloud/boat/page"
	_ "github.com/busy-cloud/iot/app"
	_ "github.com/busy-cloud/iot/internal"
	"github.com/busy-cloud/iot/protocol"
)

//go:embed pages
var pages embed.FS

//go:embed menus
var menus embed.FS

//go:embed menus
var protocols embed.FS

func init() {
	//注册页面
	page.EmbedFS(pages, "pages")

	//注册菜单
	menu.EmbedFS(menus, "menus")

	//注册协议
	protocol.EmbedFS(protocols, "protocols")
}
