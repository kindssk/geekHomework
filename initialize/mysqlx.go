package initialize

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"sync"
	"time"
)

type mysqlCon struct {
	Mysql struct {
		Database string `yaml:database`
		Host     string `yaml:host`
		Port     string `yaml:port`
		Username string `yaml:username`
		Password string `yaml:password`
		Encoding string `yaml:encoding`
		Timeout  string `yaml:timeout`
	}
}

var MysqlDB *sql.DB
var onceMysql sync.Once

func InitMysql() {
	onceMysql.Do(func() {
		tdb, err := sql.Open("mysql", mysqlConnInfo())
		if err != nil {
			panic(err)
		}
		tdb.SetConnMaxLifetime(time.Minute * 10)
		tdb.SetMaxOpenConns(10)
		tdb.SetMaxIdleConns(10)
		MysqlDB = tdb
	})
}

func mysqlConnInfo() string {
	mc := mysqlCon{}
	err := yaml.UnmarshalStrict(readConf("./conf/database.yaml"), &mc)
	if err != nil {
		panic(err)
	}
	//:111111@tcp(127.0.0.1:3306)/test?charset=utf8
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout%s", mc.Mysql.Username, mc.Mysql.Password, mc.Mysql.Host, mc.Mysql.Port, mc.Mysql.Database, mc.Mysql.Encoding, mc.Mysql.Timeout)
}
