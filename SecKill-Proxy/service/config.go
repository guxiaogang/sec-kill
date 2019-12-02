package service

import (
	"sync"
	"time"
)

const (
	ProductStatusNormal       = 0
	ProductStatusSaleOut      = 1
	ProductStatusForceSaleOut = 2
)

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Left      int
}

type ETCDConf struct {
	ETCDAddr          string
	Timeout           int
	ETCDSecKeyPrefix  string
	ETCDSecProductKey string
}

type SecKillConf struct {
	ETCDConf          ETCDConf
	RedisBlackConf    RedisConf
	RWSecProductLock  sync.RWMutex
	SecProductInfoMap map[int]*SecProductInfoConf
}
type SecResult struct {
	ProductId int
	UserId    int
	Code      int
	Token     string
}
type SecRequest struct {
	ProductId     int
	Source        string
	AuthCode      string
	SecTime       string
	Nance         string
	UserId        int
	UserAuthSign  string
	AccessTime    time.Time
	ClientAddr    string
	ClientRefence string
	CloseNotify   <-chan bool `json:"-"`

	ResultChan chan *SecResult `json:"-"`
}
