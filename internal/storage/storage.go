package storage

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"postman/internal/config"
	"postman/internal/storage/entity"
)

var lock sync.Once

type DBStorage struct {
	db *gorm.DB
}

func (s *DBStorage) GetConn() *gorm.DB {
	return s.db
}

func InitStorage(dbConfig config.DBConfig) *DBStorage {
	dBStorage := &DBStorage{}

	lock.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

		dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("DB connection failed", err)
		}

		dBStorage.db = dbConn
	})

	return dBStorage
}

func (db *DBStorage) CreateUser(user *entity.User) error {
	result := db.GetConn().Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DBStorage) FindUserByID(user *entity.User) (*entity.User, error) {
	result := db.GetConn().First(&user)
	if result.Error != nil {
		log.Print("error while search", result.Error)
		return nil, errors.New("db find error")
	}

	return user, nil
}
