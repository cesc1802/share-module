package gorm

import (
	"errors"
	"strings"

	gormdialects "github.com/cesc1802/share-module/db/gorm/dialects"
	"gorm.io/gorm"
)

type DBType int

type GetDBConn func(uri string) (*gorm.DB, error)

const (
	DbTypeMySQL DBType = iota + 1
	DbTypePostgres
	DbTypeSQLite
	DbTypeMSSQL
	DbTypeNotSupported
)

var dbTypeMap = map[DBType]GetDBConn{
	DbTypeMySQL: func(uri string) (*gorm.DB, error) {
		return gormdialects.MySqlDB(uri)
	},
	DbTypePostgres: func(uri string) (*gorm.DB, error) {
		return gormdialects.PostgresDB(uri)
	},
	DbTypeSQLite: func(uri string) (*gorm.DB, error) {
		return gormdialects.SQLiteDB(uri)
	},
	DbTypeMSSQL: func(uri string) (*gorm.DB, error) {
		return gormdialects.MSSqlDB(uri)
	},
	DbTypeNotSupported: func(uri string) (*gorm.DB, error) {
		return nil, nil
	},
}

func getDBType(dbType string) DBType {
	switch strings.ToLower(dbType) {
	case "mysql":
		return DbTypeMySQL
	case "postgres":
		return DbTypePostgres
	case "sqlite":
		return DbTypeSQLite
	case "mssql":
		return DbTypeMSSQL
	}

	return DbTypeNotSupported
}

func NewGormDB(dbType string, uri string) (dbConn *gorm.DB, err error) {
	gormDBType := getDBType(dbType)
	if gormDBType == DbTypeNotSupported {
		return nil, errors.New("database type is not supported")
	}
	return dbTypeMap[gormDBType](uri)
}
