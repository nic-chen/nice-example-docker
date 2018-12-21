package api

import (
	"github.com/nic-chen/nice"
	"../config"
)

func main() {
	n := nice.Instance(config.APP_NAME)

	n.SetDI("cache", nice.NewRedis(config.REDISHOST, config.REDISPWD, config.REDISDB, config.DBCONNOPEN, config.DBCONNIDLE));
	n.SetDI("db", nice.NewMysql(config.MYSQLHOST, config.MYSQLDB, config.MYSQLUSER, config.MYSQLPWD, config.DBCHARSET, config.DBCONNOPEN, config.DBCONNIDLE));

	Router();

	n.Run(config.HTTPBIND);
}
