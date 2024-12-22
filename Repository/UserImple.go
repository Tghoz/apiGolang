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
type TokenUser struct {
	User  models.User
	Token string
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

	tokenString, err := token.SignedString([]byte("olas papa")) //! ojo pelao
	if err != nil {
		return nil, err
	}

	tokeUser := &TokenUser{
		Token: tokenString,
		User:  user,
	}

	return tokeUser, nil

}

