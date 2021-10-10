package model

import (
	"time"

	"golang.org/x/oauth2"
)

//go:generate ../../bin/firestore-repo -p repository -o ../../repository User

// User - ユーザー
type User struct {
	ID                         string    `firestore:"-" firestore_key:"auto"`            // ID
	Name                       string    ``                                              // ニックネーム
	PrimaryEmail               string    `                                    unique:""` // メールアドレス
	Domain                     string    ``                                              // ドメイン
	IsAdmin                    bool      ``                                              // 管理者
	LINEWorksState             string    ``                                              // LINE Works State
	LINEWorksAccessToken       string    ``                                              // LINE Works Access Token
	LINEWorksRefreshToken      string    ``                                              // LINE Works Refresh Token
	GoogleCalendarEmail        string    ``                                              // Google Calendar メールアドレス
	GoogleCalendarAccessToken  string    ``                                              // Google Calendar Access Token
	GoogleCalendarRefreshToken string    ``                                              // Google Calendar Refresh Token
	GoogleCalendarTokenType    string    ``                                              // Google Calendar TokenType
	GoogleCalendarExpiry       time.Time ``                                              // Google Calendarr Expiry
	Meta
}

func (u *User) GetGoogleCalendarToken() *oauth2.Token {
	if u.GoogleCalendarAccessToken == "" {
		return nil
	}
	return &oauth2.Token{
		AccessToken:  u.GoogleCalendarAccessToken,
		RefreshToken: u.GoogleCalendarRefreshToken,
		TokenType:    u.GoogleCalendarTokenType,
		Expiry:       u.GoogleCalendarExpiry,
	}
}
