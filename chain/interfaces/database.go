package interfaces

import (

 "github.com/asdine/storm/v3"

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

func GetDB(networkType string) *storm.DB {

	db, err := storm.Open(GetPath(networkType) + "/blocks.db")
	if err != nil {
		log.Errorln(err)
		panic(err)
	}
	return db

}

func GetTxDB(networkType string) *storm.DB {
	db, err := storm.Open(GetPath(networkType) + "/tx.db")
	if err != nil {
		log.Errorln(err)
		panic(err)
	}
	return db

}

func GetWalletDB(networkType string) *storm.DB {
	db, err := storm.Open(GetPath(networkType) + "/wallet.db")
	if err != nil {
		log.Errorln(err)
		panic(err)
	}
	return db

}

func StoreBlock(networkType string, b Block) {

	var B Block
	B = b
	db := GetDB(networkType)
	defer db.Close()
	err := db.Save(&B)
	if err != nil {
		panic(err)
	}

}

func GetAllBlocks(networkType string) []*Block {
	db := GetDB(networkType)
	defer db.Close()
	var blocks []*Block
	err := db.AllByIndex("Height", &blocks)
	if err != nil {
		panic(err)
	}
	return blocks

}
