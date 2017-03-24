package repository

import "github.com/jinzhu/gorm"

type RoleLinkingRepository interface {

}

type roleLinker struct {
	db gorm.DB
}
