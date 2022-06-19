package component

import (
	"context"
	"gorm.io/gorm"
)

type UploadFileProvider interface {
	UploadFile(context.Context, string, []byte) (string, error)
}

type AppContext struct {
	GormDB             *gorm.DB
	UploadFileProvider UploadFileProvider
}

func NewAppContext(configuration Configuration) *AppContext {
	return &AppContext{GormDB: configuration.GormDB, UploadFileProvider: configuration.UploadFileProvider}
}

type Configuration struct {
	GormDB             *gorm.DB
	UploadFileProvider UploadFileProvider
}
