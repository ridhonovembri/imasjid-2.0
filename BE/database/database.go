package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// type DbInstance struct {
// 	Db *gorm.DB
// }

// var Database DbInstance
var Database *gorm.DB

func ConnectDb() error {
	var err error

	Database, err = gorm.Open(sqlite.Open("VMasjid.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
		// log.Fatal("Failed to connect to the database! \n", err)
		// os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	// Database.Logger = logger.Default.LogMode(logger.Info)

	// log.Println("Running Migrations")
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.MasjidInfo{})
	// db.AutoMigrate(&models.MasjidTest{})

	// Database = DbInstance{
	// 	Db: db,
	// }
	return nil
}
