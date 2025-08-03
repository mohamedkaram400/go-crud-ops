package middlewares

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"strings"

	"github.com/mohamedkaram400/go-crud-ops/internal/redis"
)

func RateLimiter(limit int, duration time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Build key of user IP Address
			ctx := context.Background()
			ip := strings.Split(r.RemoteAddr, ":")[0]
			key := "rate_limit:" + ip

			// INCR
			count, err := redisclient.Client.Incr(ctx, key).Result()
			redisclient.Client.Expire(ctx, key, duration)

			if err != nil {
				fmt.Println("Redis error:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if count == 1 {
				// EXPIRE
				redisclient.Client.Expire(ctx, key, duration)
			}

			fmt.Println("Rate limit key:", key, "count:", count)

			if count > int64(limit) {
				ttl, _ := redisclient.Client.TTL(ctx, key).Result()
				w.Header().Set("X-Ratelimit-Limit", strconv.Itoa(limit))
				w.Header().Set("X-Ratelimit-Remaining", "0")
				w.Header().Set("X-Ratelimit-Retry-After", strconv.Itoa(int(ttl.Seconds())))
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			w.Header().Set("X-Ratelimit-Limit", strconv.Itoa(limit))
			w.Header().Set("X-Ratelimit-Remaining", strconv.Itoa(limit-int(count)))
			next.ServeHTTP(w, r)
		})
	}
}
