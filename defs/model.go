package defs

import (
	"time"

	"github.com/GoRustNet/xurl/bit"
)

type Url struct {
	ID          string    `json:"id" form:"id" db:"id"`
	Url         string    `json:"url" form:"url" db:"url" binding:"required"`
	IsCustomize bool      `json:"is_customize" form:"is_customize" db:"is_customize"`
	UserID      int64     `json:"user_id" form:"user_id" db:"user_id" binding:"required"`
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
	ID        int64  `json:"id" form:"id" db:"id"`
	UrlID     string `json:"url_id" form:"url_id" db:"url_id"`
	IP        string `json:"ip" form:"ip" db:"ip"`
	UserAgent string `json:"user_agent" form:"user_agent" db:"user_agent"`
}
