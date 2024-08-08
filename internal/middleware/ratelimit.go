package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

const (
	limitDuration = 2 * time.Minute
	requestLimit  = 10
)

func RateLimiter(rdb *redis.Client) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := getIP(r)

			key := "rate_limit:" + ip
			val, err := rdb.Get(key).Int()
			if err != nil && err != redis.Nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if val >= requestLimit {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			if err == redis.Nil {
				rdb.Set(key, 1, limitDuration)
			} else {
				rdb.Incr(key)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getIP(r *http.Request) string {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		ip = forwarded
	}
	return ip
}
