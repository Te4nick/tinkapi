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
	return entity.User{}, nil
}
func (r *AuthSqlite) UpdateUser(user entity.User) (string, error) {
	return "", nil
}
func (r *AuthSqlite) DeleteUser(name string) (string, error) {
	return "", nil
}
