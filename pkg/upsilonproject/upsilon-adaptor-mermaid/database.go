package upsilonAdaptorMermaid;

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // load the MySQL driver.
	"log"
)

var (
	db *sql.DB
)

type Node struct {
	Identifier string;
	Karma string;
}

func GetNodes(nodeType string) []Node {
	var sql string;
	var ret []Node;

	db := DbConn();

	switch (nodeType) {
		case "drone":
			sql = "SELECT identifier, datediff(now(), lastUpdated) AS lastUpdated FROM nodes WHERE serviceType LIKE '%?(?)%' "
		case "reactor":
			sql = "SELECT identifier, datediff(now(), lastUpdated) AS lastUpdated FROM nodes WHERE serviceType LIKE '%reactor%' "
		case "custodian":
			sql = "SELECT identifier, datediff(now(), lastUpdated) AS lastUpdated FROM nodes WHERE serviceType LIKE '%custodian%' "
		case "!custodian":
			sql = "SELECT identifier, datediff(now(), lastUpdated) AS lastUpdated FROM nodes WHERE serviceType NOT LIKE '%custodian%' "
		default:
			sql = "SELECT identifier, datediff(now(), lastUpdated) AS lastUpdated FROM nodes"
	}

	fmt.Println(sql);

	cursor, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	for cursor.Next() {
		var node Node;
		var lastUpdated int;
		cursor.Scan(&node.Identifier, &lastUpdated);

		if (lastUpdated > 3) {
			node.Karma = "BAD";
		} else {
			node.Karma = "GOOD";
		}


		ret = append(ret, node)
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


