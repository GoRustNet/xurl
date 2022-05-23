package db

import (
	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/str"
	"github.com/jmoiron/sqlx"
)

func UrlList(page int, where string, params ...interface{}) (*Pagination[defs.Url], error) {
	selc := SelectBuilder("urls").Fields("id,url,is_customize,user_id,is_protected,password,dateline, visit").Order("dateline DESC").Where(where).Limit(DefaultPageSize).Offset(DefaultPageSize * page)
	return Paginate[defs.Url](selc, page, params...)
}

func UrlAdd(u *defs.Url) error {
	tx, err := Tx()
	if err != nil {
		return err
	}
	if err := UrlAddWithTx(u, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func UrlAddWithTx(u *defs.Url, tx *sqlx.Tx) error {
	u.FixFields()
	sqlStr := `INSERT INTO urls (id,url,is_customize,user_id,is_protected,password,dateline) VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT(id) DO NOTHING`
	if _, err := tx.Exec(sqlStr, u.ID, u.Url, u.IsCustomize, u.UserID, u.IsProtected, u.Password, u.Dateline); err != nil {
		return err
	}
	return nil
}

func UrlExists(u *defs.Url) (bool, error) {
	selc := SelectBuilder("urls").CountFields().Where("id=$1")
	return Exists(selc, u.ID)
}

func UrlUpdateVisit(urlID string, tx *sqlx.Tx) (int64, error) {
	sqlstr := `UPDATE urls SET visit=visit+1 WHERE id=$1 RETURNING visit`
	var visit int64
	if err := tx.Get(&visit, sqlstr, urlID); err != nil {
		return 0, err
	}
	return visit, nil
}

func GetTargetUrl(u *defs.Url, tx *sqlx.Tx) (string, error) {
	sqlstr := SelectBuilder("urls").Fields("url").Where("id=$1 AND is_del=false").Limit(1).Build()
	var targetUrl string
	if err := tx.Get(&targetUrl, sqlstr, u.ID); err != nil {
		return str.Empty, err
	}
	return targetUrl, nil
}
