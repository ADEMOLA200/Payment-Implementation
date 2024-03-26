package database

import (
	"github.com/ADEMOLA200/Payment-Implementation/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
    dsn := "root:rootroot@tcp(127.0.0.1:3306)/go_paystack?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    
    // Set the global DB variable so that other packages can access it
    DB = db

    // AutoMigrate models
    if err := DB.AutoMigrate(&model.Payment{}); err != nil {
        return err
    }

    return nil
}