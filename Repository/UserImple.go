package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	dataBase "github.com/Tghoz/apiGolang/DataBase"
	models "github.com/Tghoz/apiGolang/Model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Delete(id string) error {
	var user models.User
	if err := dataBase.Db.Unscoped().Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	if err := dataBase.Db.Unscoped().Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func FindById(id string) (*models.User, error) {

	var user models.User
	userID, err := uuid.Parse(id)
	
	if err != nil {
		return nil, err
	}

	if err := dataBase.Db.Limit(1).Where("id = ?", id).First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func FindAll() ([]models.User, error) {
	var users []models.User
	result := dataBase.Db.Order("created_at DESC").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func Create(user *models.User) (err error) {
	result := dataBase.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return
}

func Update(user *models.User, body models.User) error {

	result := dataBase.Db.Model(&user).Updates(body)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

type TokenUser struct {
	User  models.User
	Token string
}

func FindOne(email string, password string) (*TokenUser, error) {
	var user models.User

	if err := dataBase.Db.Limit(1).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, errors.New("user not found")
	}

	errP := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errP != nil && errP == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errP
	}

	expires := time.Now().Add(time.Minute * 100000)

	// Define la estructura del token

	tk := &models.Token{
		UserID: user.ID,
		Name:   user.UserName,
		Email:  user.Email,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}

	// Crea el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("olas papa")) // ojo pelao
	if err != nil {
		return nil, err
	}

	tokeUser := &TokenUser{
		Token: tokenString,
		User:  user,
	}

	return tokeUser, nil

}

func FindByEmail(email string) (*models.User, error) {
	var user models.User

	dbResult := dataBase.Db.Unscoped().Where("email = ?", email).First(&user)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("database error: %w", dbResult.Error)
	}

	return &user, nil
}
