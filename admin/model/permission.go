package model

const (
	//TbPermission 权限的表名称
	TbPermission = "permission"
)

//Permission 权限结构体
type Permission struct {
	ID    uint   `model:"id" json:"id"`
	Name  string `model:"name" json:"name"`
	Route string `model:"route" json:"route"`
}

// TbName 权限的表名称
func (perm *Permission) TbName() string { return TbPermission }

// PK 权限的主键
func (perm *Permission) PK() uint { return perm.ID }
