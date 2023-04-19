package cache

import (
	"encoding/json"
	"github.com/davveo/lemonShop/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RDS *redis.Pool

func Init() error {
	var (
		rdsCfg = config.Conf.Redis
	)
	redisConn := &redis.Pool{
		MaxIdle:     rdsCfg.MaxIdle,
		MaxActive:   rdsCfg.MaxActive,
		IdleTimeout: rdsCfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", rdsCfg.Host)
			if err != nil {
				return nil, err
			}
			if rdsCfg.Password != "" {
				if _, err := c.Do("AUTH", rdsCfg.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	RDS = redisConn

	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := RDS.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RDS.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RDS.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RDS.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(k string) error {
	conn := RDS.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+k+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
