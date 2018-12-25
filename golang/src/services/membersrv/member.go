package membersrv

import (
	"context"
	"../../dao"
	"../../proto/member"
)

type memberService struct {
}

func NewMemberService() *memberService {
	return &memberService{}
}

func (*memberService) Info(ctx context.Context, req *member.Request) (*member.Response, error) {
	d := dao.NewMemberDao();

	m, _ := d.Fetch(req.Id);

	if(len(m)<1){
		return &member.Response{Id: 0, Nickname: "", Avatar: ""}, nil
	}

	return &member.Response{Id: dao.ConvertToInt32(m["id"]), Nickname: m["nickname"].(string), Avatar: m["avatar"].(string)}, nil
}
