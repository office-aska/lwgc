package model

//go:generate ../../bin/firestore-repo -p repository -o ../../repository Session

type Session struct {
	ID     string `firestore:"-" firestore_key:"auto"`
	UserID string
	Meta
}
