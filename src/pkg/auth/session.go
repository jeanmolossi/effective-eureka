package auth

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

type Session struct {
	SessID       string        `json:"sess_id"`
	StudentID    string        `json:"sess_student_id"`
	Expiration   time.Time     `json:"sess_expiration"`
	AccessToken  *AccessToken  `json:"sess_access_token"`
	RefreshToken *RefreshToken `json:"sess_refresh_token"`
}

func NewSession(sessID, studentID string, accessToken *AccessToken, refreshToken *RefreshToken, expiration time.Time) *Session {
	return &Session{
		SessID:       sessID,
		StudentID:    studentID,
		Expiration:   expiration,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (s *Session) IsExpired() bool {
	return time.Now().UTC().Local().After(
		s.Expiration.UTC().Local(),
	)
}

func (s *Session) IsValid() bool {
	return !s.IsExpired()
}

func (s *Session) IsRefreshExpired() bool {
	return time.Now().UTC().Local().After(
		s.RefreshToken.Expiration.UTC().Local(),
	)
}

func (s *Session) Hash() string {
	return base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", s.StudentID, s.SessID)),
	)
}

func Decode(hash string) (string, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return "", "", err
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid session hash")
	}

	return parts[0], parts[1], nil
}
