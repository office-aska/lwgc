package model

//go:generate ../../bin/firestore-repo -p repository -o ../../repository -sub-collection OrganizationUser

// OrganizationUser - 組織ユーザー
type OrganizationUser struct {
	ID                        string `firestore:"-" firestore_key:"auto"` // ID
	OrganizationID            string ``                                   // 組織ID
	PrimaryEmail              string ``                                   // メールアドレス
	LINEWorksAccessToken      string ``                                   // LINE Works Access Token
	GoogleCalendarEmail       string ``                                   // Google Calendar メールアドレス
	GoogleCalendarAccessToken string ``                                   // Google Calendar Access Token
	Meta
}
