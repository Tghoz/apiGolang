package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	models "github.com/Tghoz/apiGolang/Model"
	"github.com/golang-jwt/jwt/v5"
)

func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("x-access-token")
		// Obtener el token desde la cabecera
		header = strings.TrimSpace(header)

		if header == "" {
			// Si falta el token, responder con código de error 403 Forbidden
			c.JSON(http.StatusForbidden, gin.H{"message": "Missing auth token"})
			c.Abort()
			return
		}

	

		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("olas papa"), nil // Aquí deberías implementar la lógica para obtener la clave secreta de forma segura
		})

	

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		// Agregar el token al contexto para que esté disponible para el controlador
		c.Set("user", tk)

		c.Next() // Llamar a la siguiente función en la cadena de middleware
	}
}
