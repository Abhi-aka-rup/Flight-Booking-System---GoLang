package databaseInit

import (
	"flight-system/src/entities"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() *gorm.DB {
	// Load the configuration file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./databaseInit")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	connectionString := viper.GetString("database.connectionString")
	dbName := viper.GetString("database.name")

	createDatabaseIfNotExist(connectionString, dbName)
	updateDbSchema(connectionString, dbName)

	return db
}

func createDatabaseIfNotExist(connectionString string, dbName string) {
	dsn := connectionString
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//Creating the database if it does not exist
	createDbSmt := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	isDbCreated := db.Exec(createDbSmt)
	if isDbCreated.Error != nil {
		panic(isDbCreated.Error)
	}

	//Closing the current connection
	dbSql, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	defer dbSql.Close()
}

func updateDbSchema(connectionString string, dbName string) {
	dsn := connectionString + dbName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}

	autoMigrate := db.AutoMigrate(&entities.Flight{})
	if autoMigrate != nil {
		panic(autoMigrate.Error())
	}
}
