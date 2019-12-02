package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"github.com/guxiaogang/SecKill-Proxy/service"
	"time"
)

var redisPool *redis.Pool

func initSec() (err error) {
	err = initRedis()
	if err != nil {
		logs.Error("init redis failed, err:%v", err)
		return
	}

	service.InitService(secKillConf)
	initSecProductWatcher()
	return
}

func initSecProductWatcher() {
	go watchSecProductKey(secKillConf.ETCDConf.ETCDSecProductKey)
}

func watchSecProductKey(key string) {
	//cli, err := ETCDClient.New(ETCDClient.Config{
	//	Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
	//	DialTimeout: 5 * time.Second,
	//})
	//if err != nil {
	//	logs.Error("connect etcd failed, err:", err)
	//	return
	//}
}

func initRedis() (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     secKillConf.RedisBlackConf.RedisMaxIdle,
		MaxActive:   secKillConf.RedisBlackConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisBlackConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.RedisBlackConf.RedisAddr)
		},
	}
	conn := redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
		return
	}
	return
}
