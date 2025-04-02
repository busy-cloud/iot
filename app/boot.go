package app

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/web"
	_ "github.com/busy-cloud/iot/device"
	_ "github.com/busy-cloud/iot/product"
	_ "github.com/busy-cloud/iot/project"
	_ "github.com/busy-cloud/iot/protocol"
	_ "github.com/busy-cloud/iot/space"
)

func init() {
	boot.Register("app", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"log", "web", "database"},
	})
}

func Startup() error {

	err := LoadAll()
	if err != nil {
		return err
	}

	//注册APP代理
	web.Engine().Use(Proxy)

	return nil
}

func Shutdown() error {
	apps.Range(func(name string, item *App) bool {
		if item.zipReader != nil {
			_ = item.zipReader.Close()
		}
		return true
	})
	return nil
}
