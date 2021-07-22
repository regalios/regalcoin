package interfaces

import (

	badger "github.com/timshannon/badgerhold"
	"github.com/regalios/regalcoin/config"
	log "github.com/sirupsen/logrus"

)


var localDbPath = config.ChainConfig.Database.Localpath
var testnetDbPath = config.ChainConfig.Database.Testpath
var livenetDbPath = config.ChainConfig.Database.Livepath
var dbFiles = config.ChainConfig.Database.Dbnames
var cacheSize = 1024 << 20


func GetPath(networkType string) string {
	switch networkType {
	case "live":
		return livenetDbPath
	case "test":
		return testnetDbPath
	case "local":
		return localDbPath
	default:
		return localDbPath
	}
}

func GetDB(networkType string) *badger.Store {
	options := badger.DefaultOptions
	options.Dir = GetPath(networkType)
	options.ValueDir = GetPath(networkType)
	store, err := badger.Open(options)
	if err != nil {
		log.Errorln(err)
		panic(nil)
	}
	return store

}