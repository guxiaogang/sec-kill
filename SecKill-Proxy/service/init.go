package service

var (
	secKillConf *SecKillConf
)

func InitService(serviceConf *SecKillConf) (err error) {
	secKillConf = serviceConf
	return
}
