package protocol

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/mqtt"
)

func init() {
	boot.Register("protocol", &boot.Task{
		Startup: Startup,
		Depends: []string{"web", "mqtt"},
	})
}

func Startup() error {

	mqtt.SubscribeStruct[Protocol]("iot/register/protocol", func(topic string, protocol *Protocol) {
		protocols.Store(protocol.Name, protocol)
	})
	//return Load()

	return nil
}
