package sql

import (
	"fmt"
	"reflect"
	"testing"
)

type TestStruct struct {
	ID   uint   `model:"id" json:"id"`
	Name string `model:"name" json:"name"`
}

func TestUnmashalStruct(t *testing.T) {
	s := new(TestStruct)
	c := []string{"id", "name"}

	dest, err := unmarshalStruct(s, c)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dest)
}

func TestSlice(t *testing.T) {
	var s []*TestStruct
	rt := reflect.TypeOf(s)
	v := rt.Elem()
	v = v.Elem()
	model := reflect.New(v)
	mt := model.Type()
	fmt.Printf("rt:%v\n", rt)
	fmt.Printf("mt:%v\n", mt)
}

func TestSliceItf(t *testing.T) {

}
