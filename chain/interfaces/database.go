package interfaces

import (
	leveldb "github.com/syndtr/goleveldb/leveldb"
	"path/filepath"
)


type Database struct {
	Path string `json:"-"`
	Name string `json:"-"`
	Instance *leveldb.DB
	IDB
}

const DefaultDBPathPrefix = "data"
const LocalNetDBPath = "local"
const TestNetDBPath = "testnet"
const LiveNetDBPath = "live"

var localDbPath = filepath.Join(DefaultDBPathPrefix, LocalNetDBPath)
var testnetDbPath = filepath.Join(DefaultDBPathPrefix, TestNetDBPath)
var livenetDbPath = filepath.Join(DefaultDBPathPrefix, LiveNetDBPath)
var ldb *leveldb.DB
var LDB = ldb

type IDB interface {
	GetPath(networkType string) string
	GetInstance(networkType string) *leveldb.DB
	SetInstance(networkType string) error
}

func (db *Database) GetInstance(networkType string) {

	 db.SetInstance(networkType)

	defer ldb.Close()

}

func (db *Database) SetInstance(networkType string) {

	path := db.GetPath(networkType)
	ldb, _ = leveldb.OpenFile(path, nil)

}

func (db *Database) GetPath(networkType string) string {
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
	db := new(Database)
	db.GetInstance(networkType)
}