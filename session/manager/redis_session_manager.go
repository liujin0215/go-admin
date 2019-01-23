// Package manager RedisSessionManager类
package manager

import (
	"time"

	"github.com/go-redis/redis"
)

type (
	// RedisSessionManager RedisSessionManager类的结构
	RedisSessionManager struct {
		rds          *redis.Client
		tokenManager TokenManager
		ttl          time.Duration
	}

	// RedisSession RedisSession类的结构
	RedisSession struct {
		rds   *redis.Client
		key   string
		token string
	}
)

// NewRedisSessionManager 生成RedisSessionManager
func NewRedisSessionManager(opt *redis.Options, ttl uint) *RedisSessionManager {
	return &RedisSessionManager{
		rds:          redis.NewClient(opt),
		tokenManager: new(UUIDTokenManager),
		ttl:          time.Duration(ttl) * time.Second,
	}
}

// CreateSession 创建会话
// 在创建前需要删掉已有的会话，以确保单点登录，且后登陆的会将以前登陆的挤下线
// 如果key已存在则不创建，这样处理的目的是当多点在同一时间登陆同一账号时，确保只有一个会返回成功
func (mng *RedisSessionManager) CreateSession(key string) (sess *RedisSession, err error) {
	var token string
	token, err = mng.tokenManager.NewToken()
	if err != nil {
		return
	}

	err = mng.DelSession(key)
	if err != nil {
		return
	}

	var ret bool
	// hsetnx的意义在于：同一时刻还有其他地方在登陆同一账号时，确保只有一个成功
	ret, err = mng.rds.HSetNX(key, TokenKey, token).Result()
	if err != nil {
		return
	}

	if !ret {
		return nil, ErrCreateSession
	}

	err = mng.ExpireSession(key)
	if err != nil {
		return
	}

	sess = &RedisSession{
		rds:   mng.rds,
		key:   key,
		token: token,
	}
	return
}

// GetSession 获取会话
func (mng *RedisSessionManager) GetSession(key string) (sess *RedisSession, err error) {
	sess = &RedisSession{
		rds: mng.rds,
		key: key,
	}
	err = sess.getToken()
	return
}

// DelSession 删除会话
func (mng *RedisSessionManager) DelSession(key string) (err error) {
	return mng.rds.Del(key).Err()
}

// ExpireSession 会话设置过期时间
func (mng *RedisSessionManager) ExpireSession(key string) (err error) {
	return mng.rds.Expire(key, mng.ttl).Err()
}

// Key 获取会话key
func (sess *RedisSession) Key() string {
	return sess.key
}

// Token 获取会话token
func (sess *RedisSession) Token() string {
	if len(sess.token) == 0 {
		sess.getToken()
	}
	return sess.token
}

func (sess *RedisSession) getToken() (err error) {
	sess.token, err = sess.Get(TokenKey)
	return
}

// Get 获取会话的某个key对应的value
func (sess *RedisSession) Get(key string) (value string, err error) {
	return sess.rds.HGet(sess.key, key).Result()
}

// Set 获取会话的某个key对应的value
func (sess *RedisSession) Set(key, value string) (err error) {
	return sess.rds.HSet(sess.key, key, value).Err()
}
