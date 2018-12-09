package config

import (
	"encoding/json"
	adminutils "go-admin/admin/utils"
	"go-admin/store/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Addr    string
	AdminDB *sql.DBConfig
}

var Conf *Config

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	Conf = new(Config)
	err = json.NewDecoder(file).Decode(Conf)
	if err != nil {
		return err
	}

	adminDB, err := sql.NewDB(Conf.AdminDB)
	if err != nil {
		return err
	}

	adminutils.RegisterDB(adminDB)
	return nil
}
