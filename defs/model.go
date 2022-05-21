package defs

import (
	"time"

	"github.com/GoRustNet/xurl/bit"
)

type Url struct {
	ID          string `json:"id" form:"id" db:"id"`
	Url         string `json:"url" form:"url" db:"url" binding:"required"`
	IsCustomize bool   `json:"is_customize" form:"is_customize" db:"is_customize"`
	IsDel       bool   `json:"is_del" form:"is_del" db:"is_del"`
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

type UserUrl struct {
	ID          int64     `json:"id" form:"id" db:"id"`
	UserID      int64     `json:"user_id" form:"user_id" db:"user_id" binding:"required"`
	UrlID       string    `json:"url_id" form:"url_id" db:"url_id" binding:"required"`
	IsProtected bool      `json:"is_protected" form:"is_protected" db:"is_protected"`
	Password    string    `json:"password" form:"password" db:"password"`
	IsDel       bool      `json:"is_del" form:"is_del" db:"is_del"`
	Dateline    time.Time `json:"dateline" form:"dateline" db:"dateline"`
}

type UrlVisit struct {
	ID        int64  `json:"id" form:"id" db:"id"`
	UserUrlID int64  `json:"user_url_id" form:"user_url_id" db:"user_url_id"`
	IP        string `json:"ip" form:"ip" db:"ip"`
	UserAgent string `json:"user_agent" form:"user_agent" db:"user_agent"`
}

type UserUrlView struct {
	UserID    int64  `json:"user_id" form:"user_id" db:"user_id"`
	Email     string `json:"email" form:"email" db:"email"`
	UserIsDel bool   `json:"user_is_del" form:"user_is_del" db:"user_is_del"`

	UrlID       string `json:"url_id" form:"url_id" db:"url_id"`
	TargetUrl   string `json:"target_url" form:"target_url" db:"target_url"`
	IsCustomize bool   `json:"is_customize" form:"is_customize" db:"is_customize"`
	UrlIsDel    bool   `json:"url_is_del" form:"url_is_del" db:"url_is_del"`

	UserUrlID         int64     `json:"user_url_id" form:"user_url_id" db:"user_url_id"`
	IsProtected       bool      `json:"is_protected" form:"is_protected" db:"is_protected"`
	ProtectedPassword string    `json:"protected_password" form:"protected_password" db:"protected_password"`
	Dateline          time.Time `json:"dateline" form:"dateline" db:"dateline"`
}
