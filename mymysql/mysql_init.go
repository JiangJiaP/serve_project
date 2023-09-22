package mymysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"utahw/model"
)

var (
	userName string = "myuser"
	password string = "mypassword"
	//ipAddrees string = "localhost"
	port    int    = 3306
	dbName  string = "mydatabase"
	charset string = "utf8"
)

var Db *sqlx.DB

func My_init(ipAddress string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	var err error
	Db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return err
}

// cid   route_id  multi_id
func IdCreate(data model.Data) {
	_, err := Db.Exec("insert into id_info values (?,?,?,?,?)", data.CId, data.RouteId, data.MultiId, data.Ifi, data.MacId)
	if err != nil {
		fmt.Printf("IdCreate have wrong")
	}
}

func CIdSearch(cid string) ([]model.Data, error) {
	var data []model.Data
	err := Db.Select(&data, "select * from id_info where cid = ? ", cid)
	if err != nil {
		fmt.Printf("MultiIdSearch have wrong")
	}
	return data, err
}

func RouteIdSearch(rid string) ([]model.Data, error) {
	var data []model.Data
	err := Db.Select(&data, "select * from id_info where route_id = ? ", rid)
	if err != nil {
		fmt.Printf("RouteIdSearch have wrong")
	}
	return data, err
}
