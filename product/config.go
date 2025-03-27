package product

import (
	"encoding/json"
	"errors"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/lib"
	"strings"
	"time"
	"xorm.io/xorm/schemas"
)

var configCache = lib.CacheLoader[ProductConfig]{
	Timeout: int64(time.Minute * 10),
	Loader: func(key string) (*ProductConfig, error) {
		var cfg ProductConfig
		ss := strings.Split(key, "/")
		has, err := db.Engine().ID(schemas.PK{ss[0], ss[1]}).Get(&cfg)
		if err != nil {
			return nil, err
		}
		if !has {
			return nil, errors.New("找不到")
		}
		return &cfg, nil
	},
}

func LoadConfig[T any](id, config string) (*T, error) {
	idd := id + "/" + config

	c, err := configCache.Load(idd)
	if err != nil {
		return nil, err
	}

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

	return &t, nil
}
