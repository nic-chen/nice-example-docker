package controller

import (
	"encoding/json"
	"github.com/nic-chen/nice"
	"nice-example/dao"
	"nice-example/config"
)

type member struct{}
var Member = member{}

func (member) Info(c *nice.Context) {
	id := c.ParamInt("id");

	n := nice.Instance(config.APP_NAME);
	d := dao.NewMemberDao();

	m, _ := d.Fetch(id);

	j, err := json.Marshal(1);
	n.Logger().Printf("j: %v", j);
	if err!=nil{
		n.Logger().Printf("j err: %v", err);
	}

    var jj int
    json.Unmarshal([]byte(j), &jj)
    n.Logger().Printf("jj: %v", jj);

	RenderJson(c, 0, "", m)
}