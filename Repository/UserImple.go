package repository

import (
	"errors"
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
			ExpiresAt: jwt.NewNumericDate(expires), // Usar jwt.NewNumericDate
		},
	}

	// Crea el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("olas papa"))
	if err != nil {
		return nil, err
	}

	// Aquí puedes devolver el token o hacer algo más con él
	// Por ejemplo, podrías regresar el token junto con el usuario
	// return &user, token

	tokeUser := &TokenUser{
		Token: tokenString,
		User:  user,
	}

	
	return tokeUser, nil

}
