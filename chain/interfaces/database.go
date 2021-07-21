package interfaces

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/regalios/regalcoin/config"
	log "github.com/sirupsen/logrus"
)


type Database struct {
	Path string `json:"-"`
	Name string `json:"-"`
	Instance *badger.DB
	currentBatch *badger.Txn
	IDB
}

type MemDB struct {
	Name string `json:"-"`
	Instance *badger.DB
	currentBatch *badger.Txn
	IDB
}

var localDbPath = config.ChainConfig.Database.Localpath
var testnetDbPath = config.ChainConfig.Database.Testpath
var livenetDbPath = config.ChainConfig.Database.Livepath
var dbFiles = config.ChainConfig.Database.Dbnames

var DB *badger.DB
var cacheSize = 1024 << 20
type Storage struct {}

type IDB interface {
	GetPath(networkType string) string
	GetInstance(networkType string) *badger.DB
	SetInstance(networkType string) error
}

func (s *Storage) GetInstance(networkType string) {

	path := s.GetPath(networkType)
	DB, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		log.Fatalln(err)
	}
	defer  DB.Close()

}



func (s *Storage) GetPath(networkType string) string {
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

func StartDB(networkType string) {
	db := new(Storage)
	db.GetInstance(networkType)
}