package internal

import (
	"RunnerGo-collector/internal/pkg"
	"RunnerGo-collector/internal/pkg/conf"
	"RunnerGo-collector/internal/pkg/dal/redis"
	log "RunnerGo-collector/internal/pkg/log"
)

func InitProjects(mode int) {
	conf.MustInitConf(mode)
	log.InitLogger()
	pkg.InitLocalIp()
	//es.InitEsClient(conf.Conf.ES.Host, conf.Conf.ES.Username, conf.Conf.ES.Password)
	redis.InitRedisClient(conf.Conf.ReportRedis.Address, conf.Conf.ReportRedis.Password, conf.Conf.ReportRedis.DB, conf.Conf.Redis.Address, conf.Conf.Redis.Password, conf.Conf.Redis.DB)
}
