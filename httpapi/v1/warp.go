package v1

import (
	"log"

	"github.com/GoRustNet/xurl/errs"
	"github.com/gin-gonic/gin"
)

func warp(f func(*gin.Context) *errs.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			log.Println(err.Debug())
			failed(err, c)
			c.Abort()
			return
		}
	}
}
