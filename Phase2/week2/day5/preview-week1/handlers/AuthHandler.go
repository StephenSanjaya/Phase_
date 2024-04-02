package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (as *AuthService) RegisterHandler(c *gin.Context) {

}
