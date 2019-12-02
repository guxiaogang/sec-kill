package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/guxiaogang/SecKill-Proxy/service"
)

var (
	secKillConf = &service.SecKillConf{}
)

func initConfig() (err error) {
	redisBlackAddr := beego.AppConfig.DefaultString("redis_black_addr", "127.0.0.1:6379")
	secKillConf.RedisBlackConf.RedisAddr = redisBlackAddr
	logs.Debug("read config successful, redis addr:%v", redisBlackAddr)

	if len(redisBlackAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] config is null", redisBlackAddr)
		return
	}
	return
}
