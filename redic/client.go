package redic

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// Client 客户端
type Client struct {
	pool *redis.Pool
}

// NewRedisClient 创建客户端
func NewRedisClient(addr string) *Client {
	ret := &Client{
		pool: &redis.Pool{
			MaxIdle:     3,
			MaxActive:   20,
			IdleTimeout: time.Second * 180,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", addr)
				if err != nil {
					return nil, err
				}
				_, _ = c.Do("select", 0)
				return c, nil
			},
		},
	}
	return ret
}

// Close 关闭所有链接
func (c *Client) Close() {
	c.pool.Close()
}

func (c *Client) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	rc := c.pool.Get()
	reply, err = rc.Do(commandName, args...)
	_ = rc.Close()
	return
}

// Do 操作
func (c *Client) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return c.do(commandName, args...)
}

// Set 设置值
func (c *Client) Set(key, value string) (err error) {
	_, err = c.do("SET", key, value)
	return
}

// Get 获取值
func (c *Client) Get(key string) (string, error) {
	return redis.String(c.do("GET", key))
}

// Push 队列
func (c *Client) Push(key string, value interface{}) (err error) {
	_, err = c.do("lpush", key, value)
	return
}

// Pop 默认为 list 的 brpop 操作
func (c *Client) Pop(key string, fn func(string, error)) {
	for {
		func() {
			time.Sleep(time.Millisecond * 800)
			rc := c.pool.Get()
			defer func() {
				_ = rc.Close()
			}()
			for {
				ret, err := redis.Strings(rc.Do("brpop", key, 5))
				if err == redis.ErrNil {
					return
				} else if err != nil {
					fn("", err)
					return
				}
				fn(ret[1], nil)

			}
		}()
	}
}

// PopByte 默认为 list 的 brpop 操作
func (c *Client) PopByte(key string, fn func([]byte, error)) {
	for {
		func() {
			time.Sleep(time.Millisecond * 800)
			rc := c.pool.Get()
			defer func() {
				_ = rc.Close()
			}()
			for {
				ret, err := redis.ByteSlices(rc.Do("brpop", key, 5))
				if err == redis.ErrNil {
					return
				} else if err != nil {
					fn(nil, err)
					return
				}
				fn(ret[1], nil)
			}
		}()
	}
}
