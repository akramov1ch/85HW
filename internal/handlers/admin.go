package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

type RateLimitConfig struct {
	IP       string `json:"ip"`
	Limit    int    `json:"limit"`
	Duration int    `json:"duration"` // duration in minutes
}

func SetRateLimitHandler(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var config RateLimitConfig
		if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		key := "rate_limit:" + config.IP
		err := rdb.Set(key, fmt.Sprintf("%d", config.Limit), time.Duration(config.Duration)*time.Minute).Err()
		if err != nil {
			http.Error(w, "Failed to set rate limit", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
