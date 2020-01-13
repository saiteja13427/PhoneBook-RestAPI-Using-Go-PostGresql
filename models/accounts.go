package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

//JTW token clain
type token struct {
	UserID uint
	jwt.StandardClaims
}
//Account Creation
type Account struct{
	gorm.Model //gets createdAt, updatedAt.... times
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool){
	
}