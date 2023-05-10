package rk

import (
	rkmysql "github.com/rookie-ninja/rk-db/mysql"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"myapp-go/internal/infrustructure/rk/db"
	"myapp-go/internal/infrustructure/rk/log"
)

func Init() {
	logger := rkentry.GlobalAppCtx.GetLoggerEntry("zap-logger").Logger
	log.New(logger.Sugar())
	mysqlEntry := rkmysql.GetMySqlEntry("myapp")
	myappDb := mysqlEntry.GetDB("myapp")
	db.New(myappDb)

}
