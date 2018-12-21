package srv

import (
	"../../api"
	"../../config"

	"github.com/nic-chen/nice"
	"github.com/nic-chen/nice/micro/registry"
	
	opentracing "github.com/opentracing/opentracing-go"
)

func RunApi(register registry.Registry, tracer opentracing.Tracer) {
	n := nice.Instance(config.APP_NAME)

	n.SetDI("cache", nice.NewRedis(config.REDISHOST, config.REDISPWD, config.REDISDB, config.DBCONNOPEN, config.DBCONNIDLE));
	n.SetDI("db", nice.NewMysql(config.MYSQLHOST, config.MYSQLDB, config.MYSQLUSER, config.MYSQLPWD, config.DBCHARSET, config.DBCONNOPEN, config.DBCONNIDLE));

	api.Router();

	n.Run(config.HTTPBIND);
}