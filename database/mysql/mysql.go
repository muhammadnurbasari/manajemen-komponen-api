package mysql

import (
	"database/sql"
	"fmt"
	"net/url"

	//driver mysql database
	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

//getConnectionString - preparation connection database mysql
func getConnectionString(dbHost, dbPort, dbUser, dbPass, dbName string) string {
	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", dbUser, dbPass, dbHost, dbPort, dbName)

	return desc
}

//ConnMySQL - create connection database mysql
func ConnMySQL(dbHost, dbPort, dbUser, dbPass, dbName string, maxIdle, maxConn int) (*sql.DB, error) {
	val := url.Values{}
	val.Add("loc", "Asia/Jakarta")

	desc := getConnectionString(dbHost, dbPort, dbUser, dbPass, dbName)
	dsn := fmt.Sprintf("%s&%s", desc, val.Encode())
	mysqldb, err := sql.Open(`mysql`, dsn)
	if err != nil {
		return nil, err
	}
	Conn = mysqldb

	errPing := mysqldb.Ping()
	if errPing != nil {
		return nil, errPing
	}

	mysqldb.SetMaxIdleConns(maxIdle)
	mysqldb.SetMaxOpenConns(maxConn)

	return mysqldb, nil
}
