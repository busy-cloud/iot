package protocol

import (
	"github.com/busy-cloud/boat/mqtt"
	"github.com/busy-cloud/boat/smart"
)

type Protocol struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Station     *smart.Form `json:"station,omitempty"`
	Options     *smart.Form `json:"options,omitempty"`
}

func Register(p *Protocol) error {
	tk := mqtt.Publish("iot/protocol/register", p)
	tk.Wait()
	return tk.Error()
}
