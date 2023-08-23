package handler

import (
	"github.com/wsldev/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.NewLogger("handler")
	db = config.GetSQLite()
}
