package internal

import (
	"encoding/json"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/mqtt"
	"strings"
	"time"
)

func subscribe() {
	mqtt.Subscribe("device/+/property", func(topic string, payload []byte) {
		ss := strings.Split(topic, "/")
		id := ss[1]
		var values map[string]any
		err := json.Unmarshal(payload, &values)
		if err != nil {
			log.Error(err)
			return
		}

		d := devices.Load(id)
		if d == nil {
			d = &Device{}
			has, err := db.Engine.ID(id).Get(d)
			if err != nil {
				log.Error(err)
				return
			}
			if !has {
				log.Error("device not exist")
				return
			}
			devices.Store(id, d)
		}

		//更新数据
		d.Values = values
		d.Updated = time.Now()
	})
}
