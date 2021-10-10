package model

//go:generate ../../bin/firestore-repo -p repository -o ../../repository Organization

// Organization - 組織
type Organization struct {
	ID                   string `firestore:"-" firestore_key:"auto"` // ID
	Name                 string ``                                   // 名称
	LINEWorksAppID       string ``                                   // LINE Works App ID
	LINEWorksConsumerKey string ``                                   // LINE Works Service API Consumer Key
	LINEWorksDomain      string ``                                   // LINE Works Domain
	Meta
}
