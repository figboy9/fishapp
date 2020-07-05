package infrastructure

import (
	"time"

	"github.com/ezio1119/fishapp-image/conf"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConn() (*gorm.DB, error) {
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

	dbConn, err := gorm.Open(conf.C.Db.Dbms, mysqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := dbConn.DB().Ping(); err != nil {
		return nil, err
	}

	if conf.C.Sv.Debug {
		dbConn.LogMode(true)
	}

	return dbConn, nil
}
