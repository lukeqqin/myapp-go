package rk

import (
	"context"
	"encoding/json"
	rkmysql "github.com/rookie-ninja/rk-db/mysql"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"myapp-go/internal/infrustructure/cos"
	"myapp-go/internal/infrustructure/rk/db"
	"myapp-go/internal/infrustructure/rk/log"
)

const myEntryType = "MyEntry"
const myEntryName = "myapp"

func LoadMyEntry() {
	rkentry.RegisterUserEntryRegFunc(registerMyEntriesFromConfig)
}
func Init() {
	logger := rkentry.GlobalAppCtx.GetLoggerEntry("zap-logger").Logger
	log.New(logger.Sugar())
	mysqlEntry := rkmysql.GetMySqlEntry("myapp")
	myappDb := mysqlEntry.GetDB("myapp")
	db.New(myappDb)
	myEntry := rkentry.GlobalAppCtx.GetEntry(myEntryType, myEntryName)
	myapp, _ := myEntry.(*MyEntry)
	cos.NewCosClient(myapp.MyApp.Cos)
}

// A struct which is for unmarshalled YAML
type MyEntry struct {
	MyApp struct {
		cos.Cos
	}
}

// An implementation of:
// type EntryRegFunc func(string) map[string]rkentry.Entry
func registerMyEntriesFromConfig(raw []byte) map[string]rkentry.Entry {
	res := make(map[string]rkentry.Entry)
	// 1: decode config map into boot config struct
	config := &MyEntry{}
	rkentry.UnmarshalBootYAML(raw, config)
	rkentry.GlobalAppCtx.AddEntry(config)
	res[myEntryName] = config
	return res
}

func (entry *MyEntry) Bootstrap(context.Context) {}

func (entry *MyEntry) Interrupt(context.Context) {}

func (entry *MyEntry) GetName() string {
	return myEntryName
}

func (entry *MyEntry) GetDescription() string {
	return ""
}

func (entry *MyEntry) GetType() string {
	return myEntryType
}

func (entry *MyEntry) String() string {
	bytes, _ := json.Marshal(entry)
	return string(bytes)
}
