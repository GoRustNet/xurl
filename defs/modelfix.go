package defs

import "github.com/GoRustNet/xurl/str"

type ModelFixer interface {
	FixFields()
}

func (u *User) FixFields() {
	str.FixLenRef(&u.Email, 255)
	str.FixLenRef(&u.Password, 60)
}

func (u *Url) FixFields() {
	str.FixLenRef(&u.ID, 6)
	str.FixLenRef(&u.Url, 255)
}

func (uu *UserUrl) FixFields() {
	str.FixLenRef(&uu.UrlID, 6)
	str.FixLenRef(&uu.Password, 60)
}

func (uv *UrlVisit) FixFields() {
	str.FixLenRef(&uv.IP, 45)
	str.FixLenRef(&uv.UserAgent, 255)
}

var (
	_ ModelFixer = &User{}
	_ ModelFixer = &Url{}
	_ ModelFixer = &UserUrl{}
	_ ModelFixer = &UrlVisit{}
)
