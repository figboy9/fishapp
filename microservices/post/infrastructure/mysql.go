package infrastructure

import (
	"database/sql"
	"time"

	"github.com/ezio1119/fishapp-post/conf"
	"github.com/go-sql-driver/mysql"
)

func NewMySQLDB() (*sql.DB, error) {
	mysqlConf := &mysql.Config{
		User:                 conf.C.Db.User,
		Passwd:               conf.C.Db.Pass,
		Net:                  conf.C.Db.Net,
		Addr:                 conf.C.Db.Host + ":" + conf.C.Db.Port,
		DBName:               conf.C.Db.Name,
		ParseTime:            conf.C.Db.Parsetime,
		Loc:                  time.Local,
		AllowNativePasswords: conf.C.Db.AllowNativePasswords,
	}

	db, err := sql.Open(conf.C.Db.Dbms, mysqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
