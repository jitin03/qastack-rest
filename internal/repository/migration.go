package database

import (
	"github.com/jinzhu/gorm"
	"github.com/jitin07/qastack/internal/project"
	"github.com/jitin07/qastack/internal/release"
)


func Migration(db *gorm.DB) error{
	if result :=db.AutoMigrate(&project.Project{}); result.Error !=nil{
		return result.Error
	}

	if result :=db.AutoMigrate(&release.Release{}); result.Error !=nil{
		return result.Error
	}
	return nil
}