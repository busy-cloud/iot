package boot

import (
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/iot/device"
	"github.com/busy-cloud/iot/internal"
	_ "github.com/busy-cloud/iot/product"
	_ "github.com/busy-cloud/iot/protocol"
)

func init() {
	boot.Register("iot", &boot.Task{
		Startup:  internal.Startup,
		Shutdown: nil,
		Depends:  []string{"log", "mqtt", "database"},
	})
}
