package db

import (
	"github.com/GoRustNet/xurl/defs"
)

const _ = "url_id, url, is_customize, is_protected, protected_password, dateline, url_is_del, user_id, email, user_is_del, visit"

func GetUserUrlLiteViewByUrl(u *defs.Url) (*defs.UserUrlLiteView, error) {
	selc := SelectBuilder("v_user_urls").Fields("url_id, url, is_customize, is_protected, protected_password, dateline,  user_id, email,  visit").Where("url_id=$1 AND url_is_del=false AND user_id=$2")
	var uu defs.UserUrlLiteView
	if err := db.Get(&uu, selc.Build(), u.ID, u.UserID); err != nil {
		return nil, err
	}
	return &uu, nil
}
