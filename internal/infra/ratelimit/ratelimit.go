package ratelimit

type RateLimitInterface interface {
    Rate(ip string, token string) (bool, error)
}