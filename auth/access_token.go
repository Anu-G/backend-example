package auth

import (
	"fmt"
	"time"

	"wmb-rest-api/config"
	"wmb-rest-api/model/entity"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenInterface interface {
	CreateAccessToken(cred *entity.UserCredential) (*TokenDetails, error)
	VerifyAccessToken(tokenStr string) (*AccessDetails, error)
	StoreAccessToken(userName string, tokenDetail *TokenDetails) error
	FetchAccessToken(accessDetail *AccessDetails) (string, error)
}

type token struct {
	cfg config.TokenConfig
}

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	ExpiredAt   int64
}

type AccessDetails struct {
	AccessUuid string
	UserName   string
}

func NewTokenService(cfg config.TokenConfig) TokenInterface {
	newToken := new(token)
	newToken.cfg = cfg
	return newToken
}

func (ts *token) CreateAccessToken(uc *entity.UserCredential) (*TokenDetails, error) {
	newTokenDetail := new(TokenDetails)
	now := time.Now().Local()
	end := now.Add(ts.cfg.AccessTokenLifeTime)
	newTokenDetail.ExpiredAt = end.Unix()
	newTokenDetail.AccessUuid = uuid.NewString()

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ts.cfg.ApplicationName,
			IssuedAt:  now.Unix(),
			ExpiresAt: end.Unix(),
		},
		UserName:   uc.Email,
		Email:      uc.Email,
		AccessUuid: newTokenDetail.AccessUuid,
	}

	token := jwt.NewWithClaims(ts.cfg.JwtSigningMethod, claims)
	if newToken, err := token.SignedString([]byte(ts.cfg.JwtSignatureKey)); err != nil {
		return nil, err
	} else {
		newTokenDetail.AccessToken = newToken
		return newTokenDetail, nil
	}
}

func (ts *token) VerifyAccessToken(tokenStr string) (*AccessDetails, error) {
	newAccessDetail := new(AccessDetails)
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		} else if method != ts.cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(ts.cfg.JwtSignatureKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != ts.cfg.ApplicationName {
		return nil, err
	}
	newAccessDetail.AccessUuid = claims["accessUUID"].(string)
	newAccessDetail.UserName = claims["userName"].(string)

	return newAccessDetail, nil
}

func (ts *token) StoreAccessToken(userName string, tokenDetail *TokenDetails) error {
	now := time.Now().Local()
	end := time.Unix(tokenDetail.ExpiredAt, 0)
	if err := ts.cfg.Redis.Set(tokenDetail.AccessUuid, userName, end.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}

func (ts *token) FetchAccessToken(accessDetail *AccessDetails) (string, error) {
	if accessDetail != nil {
		userId, err := ts.cfg.Redis.Get(accessDetail.AccessUuid).Result()
		if err != nil {
			return "", err
		}
		return userId, nil
	} else {
		return "", fmt.Errorf("invalid access")
	}
}
