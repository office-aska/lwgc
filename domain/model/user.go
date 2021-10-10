package model

//go:generate ../../bin/firestore-repo -p repository -o ../../repository User

// User - ユーザー
type User struct {
	ID           string `firestore:"-" firestore_key:"auto"`            // ID
	Name         string ``                                              // ニックネーム
	PrimaryEmail string `                                    unique:""` // メールアドレス
	IsAdmin      bool   ``                                              // 管理者
	Meta
}
