package defs

import (
	"time"

	"github.com/GoRustNet/xurl/bit"
	"github.com/GoRustNet/xurl/visit"
)

type Url struct {
	ID          string    `json:"id" form:"id" db:"id"`
	Url         string    `json:"url" form:"url" db:"url" binding:"required"`
	IsCustomize bool      `json:"is_customize" form:"is_customize" db:"is_customize"`
	UserID      int64     `json:"user_id" form:"user_id" db:"user_id"`
	IsProtected bool      `json:"is_protected" form:"is_protected" db:"is_protected"`
	Password    string    `json:"password" form:"password" db:"password"`
	Dateline    time.Time `json:"dateline" form:"dateline" db:"dateline"`
	IsDel       bool      `json:"is_del" form:"is_del" db:"is_del"`
	Visit       int64     `json:"visit" form:"visit" db:"visit"`
}

type UserStatus int

const (
	UserStatusPending UserStatus = iota
	UserStatusLocked
	UserStatusNormal
)

type UserPermission bit.Bit

const (
	UserPermissionSysGenerateUrl UserPermission = (1 << iota)
	UserPermissionCustomizeUrl
)
const (
	UserPermissionSysGenerateAndCustomizeUrl UserPermission = UserPermissionSysGenerateUrl | UserPermissionCustomizeUrl
)

type User struct {
	ID         int64          `json:"id" form:"id" db:"id"`
	Email      string         `json:"email" form:"email" db:"email" binding:"required"`
	Password   string         `json:"password" form:"password" db:"password" binding:"required"`
	Permission UserPermission `json:"permission" form:"permission" db:"permission"`
	Status     UserStatus     `json:"status" form:"status" db:"status"`
	Dateline   time.Time      `json:"dateline" form:"dateline" db:"dateline"`
	IsDel      bool           `json:"is_del" form:"is_del" db:"is_del"`
}

type UrlVisit struct {
	ID        int64     `json:"id" form:"id" db:"id"`
	UrlID     string    `json:"url_id" form:"url_id" db:"url_id"`
	IP        string    `json:"ip" form:"ip" db:"ip"`
	UserAgent string    `json:"user_agent" form:"user_agent" db:"user_agent"`
	Dateline  time.Time `json:"dateline" form:"dateline" db:"dateline"`
}

type UserUrlView struct {
	UrlID             string    `json:"url_id" db:"url_id"`
	Url               string    `json:"url" db:"url"`
	IsCustomize       bool      `json:"is_customize" db:"is_customize"`
	IsProtected       bool      `json:"is_protected" db:"is_protected"`
	ProtectedPassword string    `json:"-" db:"protected_password"`
	Dateline          time.Time `json:"dateline" db:"dateline"`
	UrlIsDel          bool      `json:"url_is_del" db:"url_is_del"`
	UserID            int64     `json:"user_id" db:"user_id"`
	Email             string    `json:"email" db:"email"`
	UserIsDel         bool      `json:"user_is_del" db:"user_is_del"`
	Visit             int64     `json:"visit" db:"visit"`
}
type UserUrlLiteView struct {
	UrlID             string    `json:"url_id" db:"url_id"`
	Url               string    `json:"url" db:"url"`
	IsCustomize       bool      `json:"is_customize" db:"is_customize"`
	IsProtected       bool      `json:"is_protected" db:"is_protected"`
	ProtectedPassword string    `json:"-" db:"protected_password"`
	Dateline          time.Time `json:"dateline" db:"dateline"`
	UserID            int64     `json:"user_id" db:"user_id"`
	Email             string    `json:"email" db:"email"`
	Visit             int64     `json:"visit" db:"visit"`
}

type UrlVisitToView struct {
	VisitID       int64              `json:"visit_id" db:"visit_id"`
	IP            string             `json:"-" db:"ip"`
	UserAgent     string             `json:"-" db:"user_agent"`
	VisitDateline time.Time          `json:"visit_dateline" db:"visit_dateline"`
	MaskedIP      string             `json:"ip" db:"-"`
	OS            visit.VisitOS      `json:"os" db:"-"`
	Browser       visit.VisitBrowser `json:"browser" db:"-"`
}

func (uv *UrlVisitToView) Mask() {
	uv.MaskedIP = visit.IpMask(uv.IP)
	ua := visit.ParseUserAgent(uv.UserAgent)
	uv.OS = ua.OS
	uv.Browser = ua.Browser
}

type UserUrlVisitView struct {
	UserUrlView
	UrlVisitToView
}
type UserUrlVisitLiteView struct {
	UserUrlLiteView
	UrlVisitToView
}
type UserUrlVisitLiteListView struct {
	*UserUrlLiteView
	Visits []*UrlVisitToView `json:"visits,omitempty"`
}
