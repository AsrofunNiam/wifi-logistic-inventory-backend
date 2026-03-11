package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/exception"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	domain "github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/twinj/uuid"
)

type CreateAuthFunc func(userID uint, tokenDetails *TokenDetails)

type AccessDetails struct {
	ID        uint
	FullName  string
	LegalName string
	Role      string
}

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessUUID          string
	RefreshUUID         string
	UserAgent           string
	RemoteAddress       string
	AtExpired           int64
	RefreshTokenExpired int64
}

func Auth(next func(c *gin.Context, auth *AccessDetails), roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check JWT Token
		tokenAuth, err := ExtractTokenMetadata(ExtractToken(c.Request))

		// // Check to logic auth or implementation access role
		// if tokenAuth.Role != "admin" {
		// 	helper.PanicIfError(exception.ErrUnauthorized)
		// }

		if err != nil {
			helper.PanicIfError(exception.ErrUnauthorized)
		}

		next(c, tokenAuth)
	}
}

func CreateToken(user *domain.User, userAgent, remoteAddress *string, createAuthFunc CreateAuthFunc) (*TokenDetails, error) {
	td := &TokenDetails{
		AtExpired:           time.Now().Add(time.Hour * 24 * 3).Unix(),
		AccessUUID:          uuid.NewV4().String(),
		RefreshTokenExpired: time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUUID:         uuid.NewV4().String(),
		UserAgent:           *userAgent,
		RemoteAddress:       *remoteAddress,
	}

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["exp"] = td.AtExpired
	atClaims["id"] = user.ID
	atClaims["full_name"] = user.FullName
	atClaims["legal_name"] = user.LegalName
	atClaims["role"] = user.Role

	atClaims["exp"] = td.AtExpired
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["exp"] = td.RefreshTokenExpired
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func ExtractTokenMetadata(stringToken string) (*AccessDetails, error) {
	token, err := VerifyToken(stringToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		userID := uint(claims["id"].(float64))
		return &AccessDetails{
			ID:        userID,
			FullName:  claims["full_name"].(string),
			LegalName: claims["legal_name"].(string),
			Role:      claims["role"].(string),
		}, nil
	}
	return nil, err
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
