package db

import (
	"github.com/GoRustNet/xurl/defs"
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
	sqlStr := `INSERT INTO urls (id,url,is_customize,user_id,is_protected,password,dateline) VALUES ($1,$2,$3) ON CONFLICT(id) DO NOTHING`
	if _, err := tx.Exec(sqlStr, u.ID, u.Url, u.IsCustomize, u.UserID, u.IsProtected, u.Password, u.Dateline); err != nil {
		return err
	}
	return nil
}
