package common

import (
	"github.com/go-redis/redis"
	"github.com/prometheus/common/log"
	//"time"
)
var redisCli *redis.Client

const (
	redisAddr  = ":6379"
)

func init(){

	redisCli = redis.NewClient(&redis.Options{Addr:redisAddr})
	err := redisCli.Ping().Err()
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
//bytes json unmar
func QueryByGrade(grade string) []byte{
	bs,err := redisCli.Get(grade).Bytes()
	if err != nil {
		log.Error(err)
		return nil
	}
	return bs
}

func GeneData(grade string,bs []byte) error {
	return redisCli.Set(grade,bs,0).Err()
}
