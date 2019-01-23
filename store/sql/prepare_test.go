package sql

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type TestStruct struct {
	ID      uint      `model:"id" json:"id"`
	Name    string    `model:"name" json:"name"`
	Created time.Time `model:"created" json:"created"`
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

func TestUnmashalTime(t *testing.T) {
	var s time.Time
	c := []string{"id", "name"}

	dest, err := unmarshalStruct(&s, c)
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

func TestSliceTime(t *testing.T) {
	var s []*time.Time
	rt := reflect.TypeOf(s)
	v := rt.Elem()
	v = v.Elem()
	model := reflect.New(v)
	mt := model.Type()
	fmt.Printf("rt:%v\n", rt)
	fmt.Printf("mt:%v\n", mt)
}

type TestStructEx struct {
	ID   uint      `model:"id" json:"id"`
	Name string    `model:"name" json:"name"`
	Time time.Time `model:"time" json:"time"`
}

func TestParseStruct(t *testing.T) {
	s := &TestStructEx{
		ID:   1,
		Name: "2",
		//Time: time.Now(),
	}
	fmt.Println(parseStruct(s))
}
