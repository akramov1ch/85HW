package main

import (
	"log"
	"net/http"

	"85HW/config"
	"85HW/internal/handlers"
	"85HW/internal/middleware"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	r := mux.NewRouter()

	r.Use(middleware.RateLimiter(rdb))

	r.HandleFunc("/api", handlers.ApiHandler).Methods("GET")
	r.HandleFunc("/admin/set-rate-limit", handlers.SetRateLimitHandler(rdb)).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
