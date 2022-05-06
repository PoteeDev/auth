package handlers

import (
	"github.com/PoteeDev/auth/auth"
)

// ProfileHandler struct
type profileHandler struct {
	rd auth.AuthInterface
	tk auth.TokenInterface
}

func NewProfile(rd auth.AuthInterface, tk auth.TokenInterface) *profileHandler {
	return &profileHandler{rd, tk}
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires    int64  `json:"expires_at"`
}

func (h *profileHandler) generateTokens(userId string) (error, map[string]interface{}) {
	ts, err := h.tk.CreateToken(userId)
	if err != nil {

		return err, nil
	}
	saveErr := h.rd.CreateAuth(userId, ts)
	if saveErr != nil {
		return saveErr, nil
	}
	return nil, map[string]interface{}{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
		"expires_at":    ts.AtExpires,
	}
}
