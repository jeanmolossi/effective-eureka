package auth

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

type Session struct {
	SessID       string    `json:"sess_id"`
	StudentID    string    `json:"sess_student_id"`
	Expiration   time.Time `json:"sess_expiration"`
	AccessToken  string    `json:"sess_access_token"`
	RefreshToken string    `json:"sess_refresh_token"`
}

func NewSession(sessID, studentID, accessToken, refreshToken string, expiration time.Time) *Session {
	return &Session{
		SessID:       sessID,
		StudentID:    studentID,
		Expiration:   expiration,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (s *Session) IsExpired() bool {
	return time.Now().UTC().Local().After(s.Expiration)
}

func (s *Session) IsValid() bool {
	return !s.IsExpired()
}

func (s *Session) Hash() string {
	return base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", s.StudentID, s.SessID)),
	)
}

func (s *Session) Decode(hash string) (string, string, error) {
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
