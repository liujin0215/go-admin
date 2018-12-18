package model

const (
	//TbAdminUser 管理员的表名称
	TbAdminUser = "adminuser"
)

//AdminUser 管理员结构体
type AdminUser struct {
	ID       uint   `model:"id" json:"id"`
	Name     string `model:"name" json:"name"`
	Password string `model:"password" json:"password"`
	RoleID   uint   `model:"role_id" json:"role_id"`
}
