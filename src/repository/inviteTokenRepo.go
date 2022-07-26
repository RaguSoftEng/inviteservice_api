package repository

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/RaguSoftEng/inviteservice_api/src/models"
	"gorm.io/gorm"
)

// Generate a 6 digits alphanumeric token
func buildToken() string {
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

// save token into database
func saveToken(db *gorm.DB, token string, userId uint) (*models.InviteToken, error) {

	var invToken = models.InviteToken{
		CreaterId: userId,
		IsEnabled: true,
		Token:     token,
	}

	if err := db.Create(&invToken).Error; err != nil {
		return &models.InviteToken{}, err
	}

	return &invToken, nil
}

// Responsible for Generate new Invite token
func GenerateToken(db *gorm.DB, userId uint) (string, error) {

	token := buildToken()

	res, err := saveToken(db, token, userId)
	if err != nil {
		return "", err
	}

	return res.Token, nil
}

// Validate token
func ValidateInviteToken(db *gorm.DB, token string) (bool, error) {

	var invToken models.InviteToken

	if err := db.Find(&invToken, models.InviteToken{Token: token}).Error; err != nil {
		return false, errors.New("[ ERROR ] Invalid Token!.")
	}

	if !invToken.IsEnabled {
		return false, errors.New("[ ERROR ] Invalid Token!.")
	}

	if !time.Now().Before(invToken.CreatedAt.AddDate(0, 0, 7)) {
		return false, errors.New("[ ERROR ] Token has expired!.")
	}

	return true, nil
}

// Diable existing token
func DisableToken(db *gorm.DB, token string, userId uint) (bool, error) {

	err := db.Model(&models.InviteToken{}).Where("creater_id=? AND token=?", userId, token).Update("is_enabled", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// Get all tokens for admin
func GetTokens(db *gorm.DB, userid uint) (*[]models.InviteToken, error) {

	var tokens []models.InviteToken

	if err := db.Where("creater_id=?", userid).Find(&tokens).Error; err != nil {
		return &[]models.InviteToken{}, err
	}

	if len(tokens) > 0 {
		for i := range tokens {
			err := db.Model(&models.User{}).Where("id=?", tokens[i].CreaterId).Take(&tokens[i].CreatedBy).Error
			if err != nil {
				return &[]models.InviteToken{}, err
			}
		}
	}

	return &tokens, nil
}
