package model

const (
	//TbPermission 菜单的表名称
	TbPermission = "permission"
)

//Permission 权限结构体
type Permission struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Route string `json:"route"`
}
