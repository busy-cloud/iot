package main

import (
	_ "github.com/busy-cloud/boat"
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/menu"
	"github.com/busy-cloud/boat/page"
	"github.com/busy-cloud/boat/web"
	_ "github.com/busy-cloud/connector"
	_ "github.com/busy-cloud/influxdb"
	_ "github.com/busy-cloud/iot/internal"
	_ "github.com/busy-cloud/modbus"
	//_ "github.com/busy-cloud/tsdb"
	_ "github.com/busy-cloud/user"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	//测试
	page.Dir("pages", "")

	menu.Dir("menus", "")
}

func main() {
	viper.SetConfigName("iot")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs

		//关闭web，出发
		_ = web.Shutdown()
	}()

	//安全退出
	defer boot.Shutdown()

	err := boot.Startup()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = web.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
