package db

import "testing"

func TestSelect(t *testing.T) {
	s := SelectBuilder("users").Fields("email,status").Where("is_del=$1").Order("id DESC").Limit(DefaultPageSize).Offset(0)
	sCount := s.ToCount("COUNT(id)")

	t.Log(s)
	t.Log(sCount.Build())
}
