package middleware

import (
	"labs-four/internal/infra/ratelimit"
	"net"
	"net/http"
)

func RateLimitMiddleware(limiter ratelimit.RateLimitInterface) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ip := getIP(r)
            token := r.Header.Get("API_KEY")

            allowed, err := limiter.Rate(ip, token)
            if err != nil {
                http.Error(w, "Internal rate limiter error", http.StatusInternalServerError)
                return
            }

            if !allowed {
                http.Error(w, "you have reached the maximum number of requests", http.StatusTooManyRequests)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}

func getIP(r *http.Request) string {
    ip := r.Header.Get("X-Real-IP")
    if ip == "" {
        ip = r.Header.Get("X-Forwarded-For")
    }
    if ip == "" {
        ip, _, _ = net.SplitHostPort(r.RemoteAddr)
    }
    if ip == "" {
        ip = "unknown"
    }
    return ip
}