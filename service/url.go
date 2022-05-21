package service

import (
	"github.com/GoRustNet/xurl/db"
	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/errs"
	"github.com/GoRustNet/xurl/pwd"
	"github.com/GoRustNet/xurl/shorturl"
	"github.com/GoRustNet/xurl/str"
	"github.com/jmoiron/sqlx"
)

func UrlAdd(u *defs.Url, uu *defs.UserUrl) (*defs.UserUrlView, *errs.Error) {
	tx, err := db.Tx()
	if err != nil {
		return nil, errs.DbError(err)
	}
	// 密码保护
	if str.IsNotEmpty(uu.Password) {
		hashedPassword, err := pwd.Hash(uu.Password)
		if err != nil {
			return nil, errs.BcryptError(err)
		}
		uu.Password = hashedPassword
	}
	// 自定义
	if u.IsCustomize {
		// 是否存在
		isExists, err := db.CustomizeUrlIsExistsWithTx(u, tx)
		if err != nil {
			tx.Rollback()
			return nil, errs.DbError(err)
		}
		// 如果存在，报错
		if isExists {
			tx.Rollback()
			return nil, errs.ExistsError("自定义网址已存在")
		}
		// 不存在
		if err := insertUrlWithUserUrl(u, uu, tx); err != nil {
			tx.Rollback()
			return nil, err
		}
	} else {
		// 非自定义
		urlID, err := shorturl.ShortUrl(u.Url)
		if err != nil {
			return nil, errs.ShortUrlError(err)
		}
		u.ID = urlID
		uu.UrlID = urlID
		// 如果存在，返回已存在的 urls 记录。并写入 user_urls
		urlID, urlIsExists, err := db.GetUnCustomizUrlIDIfExists(u, tx)
		if err != nil {
			tx.Rollback()
			return nil, errs.DbError(err)
		}
		if !urlIsExists {
			if err := insertUrlWithUserUrl(u, uu, tx); err != nil {
				tx.Rollback()
				return nil, err
			}
		} else {
			u.ID = urlID
			uu.UrlID = urlID
			if err := insertUrlWithUserUrl(nil, uu, tx); err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	uuv, err := db.GetUserUrlViewById(uu.ID, tx)
	if err != nil {
		tx.Rollback()
		return nil, errs.NotExistsOrDbError(err)
	}

	tx.Commit()
	return uuv, nil
}

func insertUrlWithUserUrl(u *defs.Url, uu *defs.UserUrl, tx *sqlx.Tx) *errs.Error {
	uuIsExists, err := db.UserUrlIsExists(uu, tx)
	if err != nil {
		return errs.DbError(err)
	}
	if uuIsExists {
		return errs.ExistsError("该用户已经添加过相同的网址")
	}
	if u != nil {
		if err := db.UrlAddWithTx(u, tx); err != nil {
			return errs.DbError(err)
		}
	}
	uuID, err := db.InsertUserUrl(uu, tx)
	if err != nil {
		return errs.DbError(err)
	}
	uu.ID = uuID
	return nil
}
