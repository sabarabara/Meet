package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProxyHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		proto := c.GetHeader("X-Forwarded-Proto")
		if proto != "" {
			c.Request.URL.Scheme = proto
		}

		host := c.GetHeader("X-Forwarded-Host")
		if host != "" {
			c.Request.Host = host
		}

		prefix := c.GetHeader("X-Forwarded-Prefix")

		fmt.Printf("DEBUG MID: prefix=%s, proto=%s\n", prefix, proto)

		c.Set("ProxyScheme", proto)
		c.Set("ProxyHost", host)
		c.Set("ProxyPrefix", prefix)

		c.Next()
	}
}
