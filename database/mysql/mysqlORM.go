package mysql

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
)

var ConnORM *gorm.DB

//getConnectionStringORM - preparation connection database mysql ORM
func getConnectionStringORM(dbHost, dbPort, dbUser, dbPass, dbName string) string {
	desc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", dbUser, dbPass, dbHost, dbPort, dbName)

	return desc
}

//ConnMySQLORM - create connection ORM database mysql
func ConnMySQLORM(dbHost, dbPort, dbUser, dbPass, dbName string, maxIdle, maxConn int) (*gorm.DB, error) {
	val := url.Values{}
	val.Add("loc", "Asia/Jakarta")

	desc := getConnectionStringORM(dbHost, dbPort, dbUser, dbPass, dbName)
	dsn := fmt.Sprintf("%s&%s", desc, val.Encode())
	mySqlOrm, ormErr := gorm.Open(`mysql`, dsn)
	if ormErr != nil {
		return nil, ormErr
	}
	ConnORM = mySqlOrm

	errPing := mySqlOrm.DB().Ping()
	if errPing != nil {
		return nil, errPing
	}

	mySqlOrm.DB().SetMaxIdleConns(maxIdle)
	mySqlOrm.DB().SetMaxOpenConns(maxConn)
	return mySqlOrm, nil
}
