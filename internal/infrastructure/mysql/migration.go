package mysql

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	// "tugas_akhir_example/internal/utils"

	"gorm.io/gorm"
)

// const currentfilepath = "internal/infrastructure/mysql/migration.go"

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&entity.User{},
		&entity.Alamat{},
		&entity.Toko{},
		&entity.Category{},
		&entity.Product{},
		&entity.Trx{},
		&entity.LogProduct{},
		&entity.FotoProduct{},
		&entity.DetailTrx{},
	)
	if err != nil {
		helper.Logger(helper.LoggerLevelError, "Failed Database Migrated", err.Error())
		// helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Failed Database Migrated")
	}

	// var count int64
	// if mysqlDB.Migrator().HasTable(&entity.Book{}) {
	// 	mysqlDB.Model(&entity.Book{}).Count(&count)
	// 	if count < 1 {
	// 		mysqlDB.CreateInBatches(booksSeed, len(booksSeed))
	// 	}
	// }

	helper.Logger(helper.LoggerLevelInfo, "Database Migrated", "")
	// helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelInfo, "Database Migrated")
}
