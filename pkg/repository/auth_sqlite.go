package repository

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"tinkapi/v2/pkg/repository/entity"
)

type AuthSqlite struct {
	db *gorm.DB
}

func NewAuthSqlite(db *gorm.DB) *AuthSqlite {
	return &AuthSqlite{db: db}
}

func (r *AuthSqlite) CreateUser(user entity.User) (string, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Name, result.Error
}
func (r *AuthSqlite) ReadUser(name string) (entity.User, error) {
	var foundUser entity.User
	result := r.db.Model(&entity.User{}).First(&foundUser, "name = ?", name)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return foundUser, nil
}
func (r *AuthSqlite) UpdateUser(user entity.User) (string, error) {
	return "", nil
}
func (r *AuthSqlite) DeleteUser(name string) (string, error) {
	return "", nil
}
