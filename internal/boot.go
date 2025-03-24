package internal

import (
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/iot/device"
	_ "github.com/busy-cloud/iot/product"
	_ "github.com/busy-cloud/iot/project"
	_ "github.com/busy-cloud/iot/protocol"
	_ "github.com/busy-cloud/iot/space"
)

func init() {
	boot.Register("iot", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"log", "mqtt", "database", "connector"},
	})
}
