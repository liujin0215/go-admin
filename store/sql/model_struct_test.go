package sql

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Menu struct {
	ID    uint   `model:"id" json:"id"`
	Name  string `model:"name" json:"name"`
	Route string `model:"route" json:"route"`
	Fid   uint   `model:"fid" json:"fid"`
}

func (m *Menu) TbName() string { return "TbMenu" }
func (m *Menu) PK() uint       { return m.ID }

func TestModelStruct(t *testing.T) {
	ms := &ModelStruct{Model: new(Menu)}
	//data := []byte(`{"model":{"id":1}}`)
	data, err := json.Marshal(ms)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", string(data))
}
