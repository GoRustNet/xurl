package db

import "testing"

func TestGetUrlVisit(t *testing.T) {
	testInit(t)
	uv, err := GetUrlVisit(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uv)
}
func TestUrlVisitListByUrlID(t *testing.T) {
	testInit(t)
	p, err := UrlVisitListByCondition(0, "url_id=$1 AND user_id=$2 AND url_is_del=false", "youtub", 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", p)
}
