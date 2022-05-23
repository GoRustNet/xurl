package v1

import (
	"strconv"
	"time"

	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/errs"
	"github.com/GoRustNet/xurl/service"
	"github.com/GoRustNet/xurl/str"
	"github.com/gin-gonic/gin"
)

func addUrl(c *gin.Context) *errs.Error {
	userID := getUserIDFromJwt(c)
	var u defs.Url
	if err := c.ShouldBindJSON(&u); err != nil {
		return errs.FormBindError(err)
	}
	u.UserID = userID
	u.Dateline = time.Now()
	uu, err := service.AddUrl(&u)
	if err != nil {
		return err
	}
	ok(uu, c)
	return nil
}

func getUrl(c *gin.Context) *errs.Error {
	userID := getUserIDFromJwt(c)
	urlID := c.Param("id")
	if str.IsEmpty(urlID) {
		return errs.InvalidParam()
	}
	u := &defs.Url{
		ID:     urlID,
		UserID: userID,
	}
	uu, err := service.GetUrl(u)
	if err != nil {
		return err
	}
	ok(uu, c)
	return nil
}

func visitUrl(c *gin.Context) *errs.Error {
	urlID := c.Param("id")
	if str.IsEmpty(urlID) {
		return errs.InvalidParam()
	}
	uv := &defs.UrlVisit{
		UrlID:     urlID,
		IP:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Dateline:  time.Now(),
	}
	visit, err := service.VisitUrl(uv)
	if err != nil {
		return err
	}
	ok(gin.H{
		"url":   urlID,
		"visit": visit,
	}, c)
	return nil
}

func getUrlVisit(c *gin.Context) *errs.Error {
	userID := getUserIDFromJwt(c)
	urlID := c.Param("id")
	if str.IsEmpty(urlID) {
		return errs.InvalidParam()
	}
	visitIDStr := c.Param("visitId")
	if str.IsEmpty(visitIDStr) {
		return errs.InvalidParam()
	}
	visitID, err := strconv.ParseInt(visitIDStr, 10, 64)
	if err != nil {
		return errs.ParseError(err)
	}
	uv, err1 := service.GetUrlVisit(visitID, userID, urlID)
	if err1 != nil {
		return err1
	}
	ok(uv, c)
	return nil
}

func urlVisitList(c *gin.Context) *errs.Error {
	userID := getUserIDFromJwt(c)
	urlID := c.Param("id")
	if str.IsEmpty(urlID) {
		return errs.InvalidParam()
	}
	page := 0
	p, uv, err := service.UrlVisitList(urlID, userID, page)
	if err != nil {
		return err
	}
	data := gin.H{
		"paginate": p,
		"url_data": uv,
	}
	ok(data, c)
	return nil
}
