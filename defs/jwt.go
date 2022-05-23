package defs

type JwtUserInfo struct {
	ID         int64          `json:"id"`
	Email      string         `json:"email"`
	Permission UserPermission `json:"permission"`
}
