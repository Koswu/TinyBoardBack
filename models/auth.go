package models

import (
	"boarderbackend/pkgs/logging"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Salt string `gorm:"not null" json:"salt"`
}


func CheckAuth(username, password string) bool{
	var auth Auth
	db.Where("username = ?", username).First(&auth)
	if auth.ID <= 0 {
		return false
	}
	saltPass := []byte(password + auth.Salt)
	hashPass := sha256.Sum256(saltPass)
	hashSlice := hashPass[:]
	encHash := hex.EncodeToString(hashSlice)
	if encHash != auth.Password {
		return false
	}
	return true
}

func RegisterUser(username, password string) bool{
	salt := getRandSalt()
	saltPass := []byte(password + salt)
	hashPass := sha256.Sum256(saltPass)
	hashSlice := hashPass[:]
	encHash := hex.EncodeToString(hashSlice)
	db.Create(&Auth{
		Username: username,
		Password: encHash,
		Salt: salt,
	})
	return true
}

func getRandSalt() string {
	salt := make([]byte, 20)
	_, err := rand.Read(salt)
	if err != nil {
		logging.Error("Generate Salt Failed", err)
		return ""
	}
	return hex.EncodeToString(salt)
}

func IsExistUser(username string) bool {
	var auth Auth
	db.Where("username = ?", username).First(&auth)
	return auth.ID > 0
}
