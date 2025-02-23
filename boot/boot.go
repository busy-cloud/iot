package boot

import (
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/iot/apis"
	"github.com/busy-cloud/iot/internal"
)

func init() {
	boot.Register("product", &boot.Task{
		Startup:  internal.Startup,
		Shutdown: nil,
		Depends:  []string{"log", "mqtt", "database"},
	})
}
