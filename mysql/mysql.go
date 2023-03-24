package mysql

import (
	"gorm.io/gorm"
)

type mysqlClient struct {
	mysqls map[string]*gorm.DB
}
