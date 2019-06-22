package main;

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db *sql.DB
)


func GetNodes(nodeType string) []string {
	var sql string;
	var ret []string;

	db := DbConn();

	switch (nodeType) {
		case "drone":
			sql = "SELECT identifier FROM nodes WHERE serviceType LIKE '%?(?)%' "
		case "reactor":
			sql = "SELECT identifier FROM nodes WHERE serviceType LIKE '%reactor%' "
		case "custodian":
			sql = "SELECT identifier FROM nodes WHERE serviceType LIKE '%custodian%' "
		case "!custodian":
			sql = "SELECT identifier FROM nodes WHERE serviceType NOT LIKE '%custodian%' "
		default:
			sql = "SELECT identifier FROM nodes"
	}

	fmt.Println(sql);

	cursor, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	for cursor.Next() {
		var name string;
		cursor.Scan(&name);

		ret = append(ret, name)
	}

	return ret;
}

func DbConn() (db *sql.DB) {
	conf := GetConfig();

	dbDriver := "mysql"

	connStr := conf.Database.User + ":" + conf.Database.Pass + "@tcp(" + conf.Database.Host + ":3306)/" + conf.Database.Name

	log.Println("connstr: " + connStr)

	db, err := sql.Open(dbDriver, connStr);

	if err != nil {
		panic(err.Error())
	}

	return db;
}


