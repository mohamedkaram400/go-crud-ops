package middlewares

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"net"
	"strings"

	"github.com/mohamedkaram400/go-crud-ops/internal/redis"
)

func RateLimiter(limit int, duration time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Build key of user IP Address
			ctx := context.Background()
			ip := getIP(r)
			key := "rate_limit:" + ip

			// Increment counter
			count, err := redisclient.Client.Incr(ctx, key).Result()

			if err != nil {
				fmt.Println("Redis error:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Set expiration only once
			if count == 1 {
				redisclient.Client.Expire(ctx, key, duration)
			}

			fmt.Println("Rate limit key:", key, "count:", count)

			// If over limit
			if count > int64(limit) {
				ttl, _ := redisclient.Client.TTL(ctx, key).Result()
				w.Header().Set("X-Ratelimit-Limit", strconv.Itoa(limit))
				w.Header().Set("X-Ratelimit-Remaining", "0")
				w.Header().Set("X-Ratelimit-Retry-After", strconv.Itoa(int(ttl.Seconds())))
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			// Set headers
			w.Header().Set("X-Ratelimit-Limit", strconv.Itoa(limit))
			w.Header().Set("X-Ratelimit-Remaining", strconv.Itoa(limit-int(count)))
			next.ServeHTTP(w, r)
		})
	}
}

func getIP(r *http.Request) string {
	// Check for X-Forwarded-For header (used by proxies and load balancers)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// It might contain multiple IPs — take the first
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}

	// Check for X-Real-IP header
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fallback: use RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	// Normalize IPv6 loopback ::1 to 127.0.0.1
	if ip == "::1" {
		return "127.0.0.1"
	}

	return ip
}