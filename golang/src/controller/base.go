package controller

import (
	"github.com/nic-chen/nice"
)

type Controller struct {
	Name string
}

type JsonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func New(name string) *Controller {
	ctl := &Controller{
		Name: name,
	}

	return ctl;
}

func RenderJson(c *nice.Context, code int, message string, data interface{}) {
	ret := new(JsonResp)
	ret.Code = code
	ret.Message = message
	if data != nil {
		ret.Data = data
	}
	c.JSON(200, ret)
}