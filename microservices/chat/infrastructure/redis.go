package infrastructure

import (
	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/go-redis/redis"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.C.Kvs.Host + ":" + conf.C.Kvs.Port,
		Password: conf.C.Kvs.Pass,
		DB:       conf.C.Kvs.Db,
		Network:  conf.C.Kvs.Net,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisFailoverClient() (*redis.Client, error) {
	addr := conf.C.Kvs.Sentinel.Host + ":" + conf.C.Kvs.Sentinel.Port
	c := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    conf.C.Kvs.Sentinel.MasterName,
		SentinelAddrs: []string{addr},
		Password:      conf.C.Kvs.Sentinel.Pass,
	})

	if _, err := c.Ping().Result(); err != nil {
		return nil, err
	}

	return c, nil
}
