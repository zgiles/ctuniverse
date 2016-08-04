package universedb

import (
	"github.com/garyburd/redigo/redis"
	"github.com/zgiles/ctuniverse"
)

type universedb struct {
	redispool *redis.Pool
}

func New(redispool *redis.Pool) *universedb {
	return &universedb{redispool}
}

func (local universedb) PutObject(o ctuniverse.UniverseObject) (error) {
	conn := local.redispool.Get()
	defer conn.Close()
	key := "ct:ctuniverse:objects:" + o.Uuid
	_, err := conn.Do("HMSET", redis.Args{key}.AddFlat(o)...)
	if err != nil {
		return err
	}
	return nil	
}	

func (local universedb) GetObject(k string) (o ctuniverse.UniverseObject, err error) {
	conn := local.redispool.Get()
	defer conn.Close()
	key := "ct:ctuniverse:objects:" + k
	o = ctuniverse.UniverseObject{}
	values, geterr := redis.Values(conn.Do("HGETALL", key))
	if geterr != nil {
		return o, geterr
	}
	scanerr := redis.ScanStruct(values, &o)
	if scanerr != nil {
		return o, scanerr
	}
	return o, nil
}
