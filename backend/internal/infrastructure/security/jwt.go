package security

import (
	"time"

	"bgray/taskApi/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey []byte
	expiresIn time.Duration
}

func NewJWTService(secretKey string, expiresIn time.Duration) *JWTService {
	return &JWTService{
		secretKey: []byte(secretKey),
		expiresIn: expiresIn,
	}
}

type jwtClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (j *JWTService) GenerateToken(user domain.User) (string, error) {
	claims := jwtClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

func (j *JWTService) ValidateToken(tokenString string) (*domain.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return &domain.TokenClaims{
			UserID:   claims.UserID,
			Username: claims.Username,
		}, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
