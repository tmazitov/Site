package jwt

import (
	"encoding/json"
	"site/app/internal/domain/models"
	"site/app/pkg/cache"
	"time"

	"github.com/cristalhq/jwt/v3"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type UserClains struct {
	jwt.RegisteredClaims
	username string
}

type RT struct {
	RefreshToken string
}

type helper struct {
	RTCache cache.Repository
}

func NewHelper(RTCache cache.Repository) Helper {
	return &helper{RTCache: RTCache}
}

type Helper interface {
	GenerateAccessToken(u models.User) ([]byte, error)
	UpdateRefreshToken(rt RT) ([]byte, error)
}

func (h *helper) UpdateRefreshToken(rt RT) ([]byte, error) {
	defer h.RTCache.Del([]byte(rt.RefreshToken))

	userBytes, err := h.RTCache.Get([]byte(rt.RefreshToken))
	if err != nil {
		return nil, err
	}

	var u models.User
	err = json.Unmarshal(userBytes, &u)
	if err != nil {
		return nil, err
	}

	return h.GenerateAccessToken(u)
}

func (h *helper) GenerateAccessToken(u models.User) ([]byte, error) {
	key := []byte(viper.GetString("jwt_key"))
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		return nil, err
	}

	builder := jwt.NewBuilder(signer)

	// TODO make Randomid UUID
	claims := UserClains{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        u.Register,
			Audience:  []string{"users"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		},
		username: u.Username,
	}
	token, err := builder.Build(claims)
	if err != nil {
		return nil, err
	}

	refreshTokenUuid := uuid.New()
	userBytes, _ := json.Marshal(u)
	err = h.RTCache.Set([]byte(refreshTokenUuid.String()), userBytes, 0)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(map[string]string{
		"token":         token.String(),
		"refresh_token": refreshTokenUuid.String(),
	})
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}
