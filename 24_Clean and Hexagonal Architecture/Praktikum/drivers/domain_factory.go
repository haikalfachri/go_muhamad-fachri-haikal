package drivers

import (
	userDomain "belajar-go-echo/businesses/users"
	userDB "belajar-go-echo/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}