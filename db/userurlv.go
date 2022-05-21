package db

import (
	"github.com/GoRustNet/xurl/defs"
	"github.com/jmoiron/sqlx"
)

func getUserUrlViewByCondition(tx *sqlx.Tx, where string, params ...interface{}) (*defs.UserUrlView, error) {
	selc := SelectBuilder("v_user_urls").Fields("user_id, email, user_is_del, url_id, target_url, is_customize, url_is_del, user_url_id, is_protected, protected_password, dateline").Where(where).Limit(1)
	var uuv defs.UserUrlView
	if err := tx.Get(&uuv, selc.Build(), params...); err != nil {
		return nil, err
	}
	return &uuv, nil
}

func GetUserUrlViewById(id int64, tx *sqlx.Tx) (*defs.UserUrlView, error) {
	return getUserUrlViewByCondition(tx, "user_url_id=$1", id)
}

func GetUserUrlViewByUserAndUrl(userID int64, urlID string, tx *sqlx.Tx) (*defs.UserUrlView, error) {
	return getUserUrlViewByCondition(tx, "user_id=$1 AND url_id=$2", userID, urlID)
}
