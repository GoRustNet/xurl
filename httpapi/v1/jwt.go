package v1

import (
	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/errs"
	"github.com/gin-gonic/gin"
)

func getUserFromJwt(c *gin.Context) (*defs.JwtUserInfo, *errs.Error) {
	return &defs.JwtUserInfo{
		ID:         1,
		Email:      "team@gorust.net",
		Permission: defs.UserPermissionSysGenerateAndCustomizeUrl,
	}, nil
}

func getUserIDFromJwt(c *gin.Context) int64 {
	ui, err := getUserFromJwt(c)
	if err != nil {

	}
	return ui.ID
}
