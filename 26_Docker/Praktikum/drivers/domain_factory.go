package drivers

import (
	userDomain "belajar-go-docker/businesses/users"
	userDB "belajar-go-docker/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}