package product

import (
	"encoding/json"
	"errors"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/iot/types"
	"time"
	"xorm.io/xorm/schemas"
)

type _config struct {
	load    int64
	content any
}

var configs lib.Map[_config]

func LoadConfig[T any](id, config string) (error, *T) {
	c := configs.Load(id + "/" + config)
	if c != nil {
		now := time.Now().Unix()
		// 10分钟内，不会重查
		if now-c.load < 60*10 {
			return nil, c.content.(*T)
		}
	}

	var cfg types.ProductConfig

	has, err := db.Engine().ID(schemas.PK{id, config}).Get(&cfg)
	if err != nil {
		return err, nil
	}
	if !has {
		return errors.New("缺少映射"), nil
	}

	var t T
	err = json.Unmarshal([]byte(cfg.Content), &t)
	if err != nil {
		return err, nil
	}

	return nil, &t
}
