package db

import (
	"database/sql"

	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/str"
	"github.com/jmoiron/sqlx"
)

func UrlList(page int, where string, params ...interface{}) (*Pagination[defs.Url], error) {
	selc := SelectBuilder("urls").Fields("id,url,is_customize,is_del").Order("id DESC").Where(where).Limit(DefaultPageSize).Offset(DefaultPageSize * page)
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
	sqlStr := `INSERT INTO urls (id,url,is_customize) VALUES ($1,$2,$3) ON CONFLICT(id) DO NOTHING`
	if _, err := tx.Exec(sqlStr, u.ID, u.Url, u.IsCustomize); err != nil {
		return err
	}
	return nil
}

func CustomizeUrlIsExistsWithTx(u *defs.Url, tx *sqlx.Tx) (bool, error) {
	selc := SelectBuilder("urls").Where("id=$1 AND is_customize=true").CountFields()
	return ExistsWithQueryer(tx, selc, u.ID)
}

func GetUnCustomizUrlIDIfExists(u *defs.Url, tx *sqlx.Tx) (urlID string, exists bool, err error) {
	selc := SelectBuilder("urls").Where("url=$1 AND is_customize=false").Fields("id")
	if err := tx.Get(&urlID, selc.Build(), u.Url); err != nil {
		if err == sql.ErrNoRows {
			return str.Empty, false, nil
		}
		return str.Empty, false, err
	}
	return urlID, true, nil
}
