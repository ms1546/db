package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// マイグレーション
	db.AutoMigrate(&User{})

	fmt.Println("Database connected and migrated!")
}

func createUser(db *gorm.DB, name, email string, age uint) {
	user := User{Name: name, Email: email, Age: age}
	db.Create(&user)
}

func getAllUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

func updateUser(db *gorm.DB, userID uint, name string) {
	var user User
	db.First(&user, userID)
	user.Name = name
	db.Save(&user)
}

func deleteUser(db *gorm.DB, userID uint) {
	var user User
	db.Delete(&user, userID)
}
