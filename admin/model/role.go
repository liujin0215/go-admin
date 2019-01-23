package model

const (
	//TbRole 角色的表名称
	TbRole = "role"
)

//Role 角色结构体
type Role struct {
	ID   uint   `model:"id" json:"id"`
	Name string `model:"name" json:"name"`
}

// TbName 角色的表名称
func (role *Role) TbName() string { return TbRole }

// PK 角色的主键
func (role *Role) PK() uint { return role.ID }
