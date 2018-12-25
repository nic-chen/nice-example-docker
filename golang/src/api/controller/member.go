package controller

import (
	"context"
	"encoding/json"
	"github.com/nic-chen/nice"
	"../../dao"
	"../../config"

	proto "../../proto/member"
)

type member struct{}
var Member = member{}

func (member) Info(c *nice.Context) {
	id := c.ParamInt("id");

	n := nice.Instance(config.APP_NAME);
	d := dao.NewMemberDao();

	m, _ := d.Fetch(id);

	j, err := json.Marshal(1);
	if err!=nil{
		n.Logger().Printf("j err: %v", err);
	}

    var jj int
    json.Unmarshal([]byte(j), &jj)

	if len(m)>0 {
		delete(m, "password")
		delete(m, "salt")
	}

	RenderJson(c, 0, "", m)
}

func (member) Basic(c *nice.Context) {
	id := c.ParamInt32("id");

    //服务名
	conn := newSrvDialer(config.MemberSrvName)

	n := nice.Instance(config.APP_NAME);
	
	n.Logger().Printf("connecting:%s", config.MemberSrvName);

    //grpc client
	client := proto.NewMemberClient(conn);

	req := &proto.Request{
		Id: id,
	}

	res, err := client.Info(context.Background(), req)
	if err != nil {
		n.Logger().Printf("dialer error: %v", err);
	}	

	RenderJson(c, 0, "", res)
}
