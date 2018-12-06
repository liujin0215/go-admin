package admin

type Permission struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Route string `json:"route"`
}
