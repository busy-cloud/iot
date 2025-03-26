package product

import (
	"encoding/json"
	"errors"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/lib"
	"time"
	"xorm.io/xorm/schemas"
)

var configCache = lib.Cache[ProductConfig]{
	Timeout: int64(time.Minute * 10),
}

func LoadConfig[T any](id, config string) (*T, error) {
	idd := id + "/" + config

	c, has := configCache.Load(idd)
	if has {
		//这里转来转去
		buf, err := json.Marshal(c.Content)
		if err != nil {
			return nil, err
		}

		var t T
		err = json.Unmarshal(buf, &t)
		if err != nil {
			return nil, err
		}

	}

	var cfg ProductConfig

	has, err := db.Engine().ID(schemas.PK{id, config}).Get(&cfg)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("缺少映射")
	}

	//这里转来转去
	buf, err := json.Marshal(cfg.Content)
	if err != nil {
		return nil, err
	}

	var t T
	err = json.Unmarshal(buf, &t)
	if err != nil {
		return nil, err
	}

	//缓存下来
	configCache.Store(idd, &cfg)

	return &t, nil
}
