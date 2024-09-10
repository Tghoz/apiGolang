package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Token struct {
	UserID uuid.UUID
	Name   string
	Email  string
	*jwt.RegisteredClaims
}
