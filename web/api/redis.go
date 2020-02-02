package main

import (
	"github.com/go-redis/redis"
	"github.com/prometheus/common/log"
)
var redisCli *redis.Client
const (
	redisAddr  = ":3306"
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
func queryByGrade(grade string) []byte{
	bs,err := redisCli.Get(grade).Bytes()
	if err != nil {
		log.Error(err)
		return nil
	}
	return bs
}




