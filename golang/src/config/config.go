//config

package config

import (
	//"time"
)

var (
	APP_NAME    = "nice"
	PPROFBIND   = []string{"localhost:2045"}
	HTTPBIND    = "0.0.0.0:8202"
	
	NamingAddr  = "http://etcd:2379"
	SrvName     = "member-srv"
	SrvHost     = "localhost"
	SrvPort     = "50001"
	SrvCheckTTL = 30
	CliName     = "member-cli"

	JaegerAddr  = "jaeger:6831"

	MemberSrvName = SrvName
)

const (
	Debug      = true
	
	MYSQLHOST  = "mysql"
	MYSQLDB    = "test"
	MYSQLUSER  = "root"
	MYSQLPWD   = "123456"
	
	DBCHARSET  = "utf8"
	DBCONNOPEN = 100
	DBCONNIDLE = 10
	
	REDISHOST  = "redis:6379"
	REDISDB    = 0
	REDISPWD   = ""

)