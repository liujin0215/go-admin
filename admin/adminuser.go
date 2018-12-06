package admin

type AdminUser struct {
	ID   uint   `model:"id" json:"id"`
	Name string `model:"name" json:"name"`
}
