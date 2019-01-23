package sql

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFindOne(t *testing.T) {
	cfg := &DBConfig{
		Addr: "root:Liujin@tcp(localhost:3306)/data?charset=utf8mb4&parseTime=true",
		Type: "mysql",
	}

	db, err := NewDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	s := &TestStruct{ID: 1}

	err = db.Tb("menu").Select(s).Where(s).FindOne(s).Err()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(s)
}

func TestFind(t *testing.T) {
	cfg := &DBConfig{
		Addr: "root:Liujin@tcp(localhost:3306)/data?charset=utf8mb4&parseTime=true",
		Type: "mysql",
	}

	db, err := NewDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	s := &TestStruct{ID: 1}
	//var records time.Time

	err = db.Tb("adminuser").Select("created").Where("id = ?", 1).FindOne(s).Err()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(s)
}
