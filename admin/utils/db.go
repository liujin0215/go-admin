package utils

import (
	"go-admin/store/sql"
)

var (
	//AdminDB 管理员DB
	AdminDB *sql.DB
)

//RegisterDB 注册管理员DB
func RegisterDB(db *sql.DB) {
	AdminDB = db
}
