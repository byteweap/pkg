package redix

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/byteweap/pkg/convert"
	"github.com/byteweap/pkg/logs"
)

type Redix struct {
	rdb  redis.UniversalClient
	ctx  context.Context
	stop chan struct{} // 终止信号
}

func New(addr, password string) (*Redix, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Redix{rdb: rdb, ctx: ctx, stop: make(chan struct{})}, nil
}

func (r *Redix) Get(key string) (string, error) {
	val, err := r.rdb.Get(r.ctx, key).Result()
	if errors.Is(err, redis.Nil) { // key不存在
		return val, nil
	}
	return val, err
}

func (r *Redix) GetInt64(key string) (int64, error) {
	data, err := r.rdb.Get(r.ctx, key).Int64()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	return data, err
}

func (r *Redix) GetBool(key string) (bool, error) {
	data, err := r.rdb.Get(r.ctx, key).Bool()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	return data, err
}

func (r *Redix) Set(key string, val interface{}, dur time.Duration) error {
	return r.rdb.Set(r.ctx, key, val, dur).Err()
}

func (r *Redix) Del(key string) error {
	return r.rdb.Del(r.ctx, key).Err()
}

func (r *Redix) LPush(key string, data []byte) error {
	return r.rdb.LPush(r.ctx, key, data).Err()
}

func (r *Redix) ListenAppMsg(appName string, fn func(data []byte)) {

	// 监听前清除队列内的数据
	_ = r.Del(appName)

	go func() {
		for {
			select {
			case <-r.stop:
				logs.Infox().Msgf("redix stop!!!!!")
				return
			default:
				val := r.rdb.BRPop(r.ctx, time.Second, appName).Val()
				//fmt.Println("接收到List数据: ", msg, " err: ", err)
				// val[0]: key值
				// val[1]: 传输的值
				if len(val) == 2 {
					fn(convert.String2Bytes(val[1]))
				}
			}
		}
	}()
}

func (r *Redix) Incrby(key string, incr int64) error {
	return r.rdb.IncrBy(r.ctx, key, incr).Err()
}

func (r *Redix) HSetMap(key string, val map[string]interface{}) error {
	return r.rdb.HSet(r.ctx, key, val).Err()
}

func (r *Redix) HGet(key, field string) {
	val := r.rdb.HGet(r.ctx, key, field).Val()
	fmt.Println("HGet val: ", val)
}

func (r *Redix) HGetAll(key string) (map[string]string, error) {

	data, err := r.rdb.HGetAll(r.ctx, key).Result()
	if errors.Is(err, redis.Nil) { // key不存在
		return map[string]string{}, nil
	}
	return data, err
}

func (r *Redix) HDel(key string, fields ...string) error {
	_, err := r.rdb.HDel(r.ctx, key, fields...).Result()
	if err != nil {
		return err
	}
	return nil
}

// HIncrBy 增加key下field的值(int)
// 线程安全(已测)
func (r *Redix) HIncrBy(key, field string, incr int64) {
	num, err := r.rdb.HIncrBy(r.ctx, key, field, incr).Result()
	fmt.Println("num: ", num, " err: ", err)
}

// SetEx 设置过期时间
func (r *Redix) SetEx(key string, expiration time.Duration) {
	ret, err := r.rdb.SetEx(r.ctx, key, nil, expiration).Result()
	fmt.Println("SetEx ret: ", ret, " err: ", err)
}

func (r *Redix) Close() {
	close(r.stop) // 终止信号
	_ = r.rdb.Close()
}
