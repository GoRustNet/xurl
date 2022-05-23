package db

import (
	"github.com/GoRustNet/xurl/defs"
	"github.com/jmoiron/sqlx"
)

// "url_id, url, is_customize, is_protected, protected_password, dateline, url_is_del, user_id, email, user_is_del, visit, visit_id, ip, user_agent, visit_dateline"

func InsertUrlVisit(uv *defs.UrlVisit, tx *sqlx.Tx) (int64, error) {
	sqlstr := `INSERT INTO url_visits (url_id, ip, user_agent, dateline) VALUES ($1,$2,$3,$4) RETURNING id`
	var id int64
	if err := tx.Get(&id, sqlstr, uv.UrlID, uv.IP, uv.UserAgent, uv.Dateline); err != nil {
		return 0, err
	}
	return id, nil
}

func GetUrlVisit(visitID int64, userID int64, urlID string) (*defs.UserUrlVisitLiteView, error) {
	selc := SelectBuilder("v_user_url_visits").Fields("url_id, url, is_customize, is_protected, protected_password, dateline,  user_id, email,  visit, visit_id, ip, user_agent, visit_dateline").Where("visit_id=$1 AND user_id=$2 AND url_id=$3 AND url_is_del=false").Limit(1)
	sqlStr := selc.Build()
	var uuv defs.UserUrlVisitLiteView
	if err := db.Get(&uuv, sqlStr, visitID, userID, urlID); err != nil {
		return nil, err
	}
	return &uuv, nil
}

func UrlVisitListByCondition(page int, condition string, params ...interface{}) (*Pagination[defs.UserUrlVisitLiteView], error) {
	selc := SelectBuilder("v_user_url_visits").Fields("url_id, url, is_customize, is_protected, protected_password, dateline,  user_id, email,  visit, visit_id, ip, user_agent, visit_dateline").Where(condition).Limit(DefaultPageSize).Offset(page * DefaultPageSize).Order("visit_id DESC")
	return Paginate[defs.UserUrlVisitLiteView](selc, page, params...)
}

func UrlVisitListByUrl(page int, urlID string, userID int64) (*Pagination[defs.UserUrlVisitLiteView], error) {
	return UrlVisitListByCondition(page, "url_id=$1 AND user_id=$2", urlID, userID)
}
