package database

import (
	"github.com/ADEMOLA200/Payment-Implementation/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

    if dbUser == "" || dbPassword == "" {
        panic("DB_USER or DB_PASSWORD environment variables are not set")
    }

    dsn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("could not connect to the database: " + err.Error())
    }

    if err := DB.AutoMigrate(
	    &model.Payment{}
    ); err != nil {
        return err
    }

    return nil
}
