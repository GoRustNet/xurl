package db

import (
	"github.com/GoRustNet/xurl/defs"
	"github.com/jmoiron/sqlx"
)

func InsertUserUrl(uu *defs.UserUrl, tx *sqlx.Tx) (int64, error) {
	uu.FixFields()
	sqlStr := `INSERT INTO user_urls (user_id,url_id,is_protected,password,dateline) VALUES($1,$2,$3,$4,$5) RETURNING id`
	var id int64
	if err := tx.Get(&id, sqlStr, uu.UserID, uu.UrlID, uu.IsProtected, uu.Password, uu.Dateline); err != nil {
		return 0, err
	}
	return id, nil
}

func UserUrlIsExists(uu *defs.UserUrl, tx *sqlx.Tx) (bool, error) {
	selc := SelectBuilder("user_urls").Where("user_id=$1 AND url_id=$2").CountFields()
	return ExistsWithQueryer(tx, selc, uu.UserID, uu.UrlID)
}
