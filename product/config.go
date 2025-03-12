package product

import (
	"encoding/json"
	"errors"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/lib"
	"time"
	"xorm.io/xorm/schemas"
)

type _config struct {
	load    int64
	content any
}

var cache lib.Map[_config]

func LoadConfig[T any](id, config string) (error, *T) {
	idd := id + "/" + config

	c := cache.Load(idd)
	if c != nil {
		now := time.Now().Unix()
		// 10分钟内，不会重查
		if now-c.load < 60*10 {
			return nil, c.content.(*T)
		}
	}

	var cfg ProductConfig

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

	//缓存下来
	cache.Store(idd, &_config{
		load:    time.Now().Unix(),
		content: &t,
	})

	return nil, &t
}
