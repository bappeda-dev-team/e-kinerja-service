package middle_ware

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"

	"github.com/labstack/echo/v4"
)

type ipLimiter struct {
	limiter *rate.Limiter
}

var (
	authLimiters = make(map[string]*ipLimiter)
	authMu       sync.Mutex
)

func getAuthLimiter(ip string) *rate.Limiter {
	authMu.Lock()
	defer authMu.Unlock()

	if l, ok := authLimiters[ip]; ok {
		return l.limiter
	}

	// 5 request per menit, burst maksimal 5
	l := &ipLimiter{limiter: rate.NewLimiter(rate.Limit(5.0/60.0), 5)}
	authLimiters[ip] = l
	return l.limiter
}

// AuthRateLimiter membatasi request ke auth endpoints: 5 request/menit per IP
func AuthRateLimiter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		if !getAuthLimiter(ip).Allow() {
			return c.JSON(http.StatusTooManyRequests, map[string]interface{}{
				"code":    429,
				"message": "Terlalu banyak percobaan, coba lagi dalam beberapa menit",
			})
		}
		return next(c)
	}
}
