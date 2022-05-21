package db

import "github.com/GoRustNet/xurl/defs"

func UserList(page int) (*Pagination[defs.User], error) {
	selc := SelectBuilder("users").Fields("id,email,password,permission,status,dateline,is_del").Order("id DESC")
	return Paginate[defs.User](selc, page)
}

func UserAdd(u *defs.User) (int64, error) {
	u.FixFields()
	var userID int64
	sql := `INSERT INTO users (email, password, permission, status, dateline) VALUES($1,$2,$3,$4,$5)  RETURNING id`
	if err := db.Get(&userID, sql, u.Email, u.Password, u.Permission, u.Status, u.Dateline); err != nil {
		return 0, err
	}
	return userID, nil
}

func UserExists(u *defs.User) (bool, error) {
	u.FixFields()
	selc := SelectBuilder("users").CountFields().Where("email=$1")
	return Exists(selc, u.Email)
}

func GetUser(where string, params ...interface{}) (*defs.User, error) {
	selc := SelectBuilder("users").Fields("id,email,password,permission,status,dateline,is_del").Where(where).Limit(1)
	var u defs.User
	if err := db.Get(&u, selc.Build(), params...); err != nil {
		return nil, err
	}
	return &u, nil
}
