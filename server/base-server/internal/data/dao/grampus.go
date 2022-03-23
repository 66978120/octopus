package dao

import (
	"gorm.io/gorm"
)

type GrampusDao interface {
}

type grampusDao struct {
	db *gorm.DB
}

func NewGrampusDao(db *gorm.DB) GrampusDao {
	return &grampusDao{
		db: db,
	}
}
