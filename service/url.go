package service

import (
	"database/sql"
	"fmt"

	"github.com/GoRustNet/xurl/db"
	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/errs"
	"github.com/GoRustNet/xurl/shorturl"
	"github.com/GoRustNet/xurl/str"
)

func AddUrl(u *defs.Url) (*defs.UserUrlLiteView, *errs.Error) {
	if u.IsCustomize {
		if str.IsEmpty(u.ID) {
			return nil, errs.InvalidParam("请指定自定义URL")
		}
	} else {
		urlID, err := shorturl.ShortUrlWithSeed(u.Url, uint32(u.UserID))
		if err != nil {
			return nil, errs.ShortUrlError(err)
		}
		u.ID = urlID
	}
	u.FixFields()
	urlIsExists, err := db.UrlExists(u)
	if err != nil {
		return nil, errs.DbError(err)
	}
	if urlIsExists {
		return nil, errs.ExistsError(fmt.Sprintf("短网址'%s'已存在", u.ID))
	}
	if err := db.UrlAdd(u); err != nil {
		return nil, errs.DbError(err)
	}

	uu, err := db.GetUserUrlLiteViewByUrl(u)
	if err != nil {
		return nil, errs.NotExistsOrDbError(err)
	}

	return uu, nil
}

func GetUrl(u *defs.Url) (*defs.UserUrlLiteView, *errs.Error) {
	uu, err := db.GetUserUrlLiteViewByUrl(u)
	if err == sql.ErrNoRows {
		return nil, errs.NotExistsError("网址不存在")
	}
	if err != nil {
		return nil, errs.DbError(err)
	}
	return uu, nil
}

func VisitUrl(uv *defs.UrlVisit) (int64, *errs.Error) {
	tx, err := db.Tx()
	if err != nil {
		return 0, errs.DbError(err)
	}
	visit, err := db.UrlUpdateVisit(uv.UrlID, tx)
	if err != nil {
		tx.Rollback()
		return 0, errs.NotExistsOrDbError(err)
	}
	if _, err := db.InsertUrlVisit(uv, tx); err != nil {
		tx.Rollback()
		return 0, errs.NotExistsOrDbError(err)
	}
	tx.Commit()
	return visit, nil
}

func GetUrlVisit(visitID, userID int64, urlID string) (*defs.UserUrlVisitLiteView, *errs.Error) {
	uv, err := db.GetUrlVisit(visitID, userID, urlID)
	if err != nil {
		return nil, errs.NotExistsOrDbError(err)
	}
	uv.Mask()
	return uv, nil
}
func UrlVisitList(urlID string, userID int64, page int) (*db.Pagination[defs.UserUrlVisitLiteView], *defs.UserUrlVisitLiteListView, *errs.Error) {
	p, err := db.UrlVisitListByCondition(page, "url_id=$1 AND user_id=$2 AND url_is_del=false", urlID, userID)
	if err != nil {
		return nil, nil, errs.NotExistsOrDbError(err)
	}
	v := new(defs.UserUrlVisitLiteListView)
	if p.Data != nil && len(p.Data) > 0 {
		v.Visits = make([]*defs.UrlVisitToView, 0, len(p.Data))
		for _, vv := range p.Data {
			uv2v := &vv.UrlVisitToView
			uv2v.Mask()
			v.Visits = append(v.Visits, uv2v)
		}
		v.UserUrlLiteView = &(p.Data[0].UserUrlLiteView)
	}
	p.Data = nil
	return p, v, nil
}
