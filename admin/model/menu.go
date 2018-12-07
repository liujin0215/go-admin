package model

const (
	//TbMenu 菜单的表名称
	TbMenu = "menu"
)

//Menu 菜单结构体
type Menu struct {
	ID   uint   `model:"id" json:"id"`
	Name string `model:"name" json:"name"`
	Fid  uint   `model:"fid" json:"fid"`
}
