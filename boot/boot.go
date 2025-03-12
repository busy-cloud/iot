package boot

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/iot/internal"
)

func init() {
	boot.Register("iot", &boot.Task{
		Startup:  internal.Startup,
		Shutdown: nil,
		Depends:  []string{"log", "mqtt", "database"},
	})
}
