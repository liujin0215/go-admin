package model

const (
	//TbRole 菜单的表名称
	TbRole = "role"
)

//Role 角色结构体
type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
