package Model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetUserByUsername(username string,db *gorm.DB) (*User, error) {
	var user User
	err:= db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	if err.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (u *User) Save(db *gorm.DB) (error){
	err := u.BeforeSave(db)
	if err != nil {
		return err
	}

	err = db.Exec("INSERT INTO users (username, password, created_at, updated_at) VALUES (?, ?, ?, ?)", u.Username, u.Password, time.Now(),time.Now()).Error
	if err != nil {
		return errors.New("user already exists")
	}
	return nil
}

func (u *User) BeforeSave(db *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

