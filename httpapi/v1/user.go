package v1

import (
	"github.com/GoRustNet/xurl/errs"
	"github.com/gin-gonic/gin"
)

func userRegister(c *gin.Context) *errs.Error {
	ok(gin.H{}, c)
	return nil
}
