package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"restexample/models"
)

type UserRepository interface {
	Save(user models.User) (uint, error)
	GetAllUsers() []*models.User
}
type userDatabase struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	if DB == nil {
		_, err = Connect()
		if err != nil {
			log.Error("DB connection issue")
		}
	}

	return &userDatabase{
		connection: DB,
	}
}

func (db userDatabase) GetAllUsers() []*models.User {
	var users []*models.User
	db.connection.Find(&users)
	return users
}

func (db userDatabase) Save(user models.User) (uint, error) {
	result := db.connection.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}
