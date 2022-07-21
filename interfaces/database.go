package interfaces

import "gorm.io/gorm"

type SQL interface {
	Orm() *gorm.DB
	MySQL() error
}
