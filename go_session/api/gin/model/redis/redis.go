package model_redis

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var conn *redis.Client

// 初期化関数：Redisとの接続
func init() {
	redisContainerIPv4 := os.Getenv("REDIS_CONTAINER_IPV4") // RedisコンテナIPv4
	redisContainerPort := os.Getenv("REDIS_CONTAINER_PORT") // Redisコンテナポート番号
	redisURL := redisContainerIPv4 + ":" + redisContainerPort
	conn = redis.NewClient(&redis.Options{
		Addr:     redisURL, // Redis URL
		Password: "",       // no password set
		DB:       0,        // use default DB
	})
}

// 新しいセッションの作成関数
func NewSession(c *gin.Context, cookieKey, redisValue string) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("ランダムな文字作成時にエラーが発生しました。")
	}
	newRedisKey := base64.URLEncoding.EncodeToString(b)

	if err := conn.Set(c, newRedisKey, redisValue, 0).Err(); err != nil {
		panic("Session登録時にエラーが発生：" + err.Error())
	}
	c.SetCookie(cookieKey, newRedisKey, 0, "/", "localhost", false, false)
}

// セッションの取得関数
func GetSession(c *gin.Context, cookieKey string) interface{} {
	redisKey, _ := c.Cookie(cookieKey)
	redisValue, err := conn.Get(c, redisKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("SessionKeyが登録されていません。")
		return nil
	case err != nil:
		fmt.Println("Session取得時にエラー発生：" + err.Error())
		return nil
	}
	return redisValue
}

// セッションの削除関数
func DeleteSession(c *gin.Context, cookieKey string) {
	redisId, _ := c.Cookie(cookieKey)
	conn.Del(c, redisId)
	c.SetCookie(cookieKey, "", -1, "/", "localhost", false, false)
}
