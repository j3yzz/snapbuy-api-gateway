package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

const (
	BUCKET    = 1 * 60
	EXPIRY    = 5 * 60
	THRESHOLD = 2
)

func RateLimitMiddleware(cache *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		IPAddress := c.Request.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = c.Request.Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress = c.Request.RemoteAddr
		}
		IPAddress = getKey(IPAddress)
		log.Println("IP:", IPAddress)
		key := "rate-limit-" + IPAddress + "-counter"
		getRateLimit, err := cache.Get(key).Result()
		fmt.Println("types:", reflect.TypeOf(getRateLimit), reflect.TypeOf(THRESHOLD))
		if err != nil {
			getRateLimit = cache.Set(key, 1, 0).Val()
		} else {
			getRateLimitInteger, _ := strconv.Atoi(getRateLimit)
			if getRateLimitInteger > THRESHOLD {

				c.JSON(http.StatusBadGateway, gin.H{
					"success": false,
					"message": "Max rate limiting reached, please try after some time",
				})
				c.Abort()
				return
			}
			getRateLimit = strconv.FormatInt(cache.Incr(key).Val(), 10)
		}
		log.Println(key, getRateLimit)

		c.Next()
	}
}

func getKey(IP string) string {
	bucket := time.Now().Unix() / BUCKET
	IP = IP + strconv.FormatInt(bucket, 10)
	return IP
}
