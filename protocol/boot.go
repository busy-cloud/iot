package protocol

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/mqtt"
	"strings"
)

func init() {
	boot.Register("protocol", &boot.Task{
		Startup: Startup,
		Depends: []string{"web", "mqtt"},
	})
}

func Startup() error {

	mqtt.SubscribeStruct[Protocol]("register/protocol/+", func(topic string, protocol *Protocol) {
		name := strings.TrimPrefix(topic, "register/protocol/")
		protocols.Store(name, protocol)
	})
	//return Load()

	return nil
}
