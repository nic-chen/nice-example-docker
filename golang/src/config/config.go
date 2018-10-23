//config
package config

import (
	//"time"
)

var (
	APP_NAME  = "nice";
	HTTPBIND  = "0.0.0.0:8090"
)

const (
	Debug      = true;
	
	MYSQLHOST  = "mysql";
	MYSQLDB    = "test";
	MYSQLUSER  = "root";
	MYSQLPWD   = "";
	
	DBCHARSET  = "utf8";
	DBCONNOPEN = 100;
	DBCONNIDLE = 10;
	
	REDISHOST  = "redis:6379";
	REDISDB    = 0;
	REDISPWD   = "";

)