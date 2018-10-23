package main

import (
	"github.com/nic-chen/nice"
	"nice-example/config"
	"nice-example/controller"
)

func Router() {
	n := nice.Instance(config.APP_NAME)
	
	n.Get("/member/:id", controller.Member.Info);
}