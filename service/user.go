package service

import (
	"github.com/GoRustNet/xurl/db"
	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/errs"
	"github.com/GoRustNet/xurl/pwd"
)

func UserRegister(u *defs.User) *errs.Error {
	hashedPassword, err := pwd.Hash(u.Password)
	if err != nil {
		return errs.BcryptError(err)
	}
	u.Password = hashedPassword
	isExists, err := db.UserExists(u)
	if err != nil {
		return errs.DbError(err)
	}
	if isExists {
		return errs.ExistsError("用户已存在")
	}
	id, err := db.UserAdd(u)
	if err != nil {
		return errs.DbError(err)
	}
	u.ID = id
	return nil
}

func getUser(where string, params ...interface{}) (*defs.User, *errs.Error) {
	u, err := db.GetUser(where, params...)
	if err != nil {
		return nil, errs.NotExistsOrDbError(err, "不存在的用户")
	}
	return u, nil
}

func GetUserByEmail(email string) (*defs.User, *errs.Error) {
	return getUser("email=$1", email)
}
func GetUserById(id int64) (*defs.User, *errs.Error) {
	return getUser("id=$1", id)
}
