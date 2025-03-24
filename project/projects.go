package project

import (
	"fmt"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/iot/space"
)

var projects lib.Map[Project]

func Get(id string) *Project {
	return projects.Load(id)
}

func Load(id string) error {
	var m Project
	has, err := db.Engine().ID(id).Get(&m)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("找不到项目%s", id)
	}

	return From(&m)
}

func From(p *Project) error {
	projects.Store(p.Id, p)

	var ds []*space.Space
	err := db.Engine().Where("project_id=?", p.Id).Find(&ds)
	if err != nil {
		return err
	}

	for _, s := range ds {
		err := space.From(s)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}
