package controllers

import (
	"context"
	"crypto/rand"
	"time"

	"net/http"
	"os"

	"encoding/base64"
	"encoding/json"
	"io"

	models "github.com/Tghoz/apiGolang/Model"
	repo "github.com/Tghoz/apiGolang/Repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func init() {
	godotenv.Load("google.env")

	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
}

func generateRandomPassword(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func GoogleLogin(c *gin.Context) {

	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)

}

func GoogleRedirect(c *gin.Context) {

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code is empty"})
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	defer response.Body.Close()
	var googleUser map[string]interface{}
	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(bytes, &googleUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Create a User struct with the data from Google
	email := googleUser["email"].(string)
	user, err := repo.FindByEmail(email)

	if err != nil {
		if err.Error() == "user not found" {
			randomPassword, err := generateRandomPassword(12)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating random password: " + err.Error()})
				return
			}
			user = &models.User{
				ID:       uuid.New(),
				Email:    email,
				UserName: googleUser["name"].(string),
				Password: randomPassword,
			}
			err = repo.Create(user)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user: " + err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

	expires := time.Now().Add(time.Minute * 100000)
	tk := &models.Token{
		UserID: user.ID,
		Name:   user.UserName,
		Email:  user.Email,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, err := tokenJwt.SignedString([]byte("olas papa")) // ojo pelao

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verificar si tokenString está vacío
	if tokenString == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Generated token string is empty"})
		return
	}

	c.Header("Access-Control-Expose-Headers", "x-access-token")
	c.Header("x-access-token", tokenString)

	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"user":          user,
		"session_token": tokenString,
	})

}
