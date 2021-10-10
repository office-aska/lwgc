package environ

import (
	"os"
	"strings"

	"github.com/office-aska/lwgc/pkg/definitions"
)

var (
	projectID         = getProjectID()
	serviceName       = os.Getenv(definitions.EnvKeyServiceName)
	serviceVersion    = os.Getenv(definitions.EnvKeyServiceVersion)
	ServiceHostDomain string
)

func getProjectID() string {
	id, ok := os.LookupEnv(definitions.EnvKeyProjectID)
	if ok {
		return id
	}

	id, ok = os.LookupEnv(definitions.EnvKeyGoogleCloudProject)
	if ok {
		return id
	}

	return ""
}

// IsLocal - ローカル環境か判断する
func IsLocal() bool {
	return os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") != ""
}

// IsDev - dev環境か判断する
func IsDev() bool {
	return strings.HasSuffix(projectID, "-dev")
}

// IsQA - qa環境か判断する
func IsQA() bool {
	return strings.HasSuffix(projectID, "-qa")
}

// IsStaging - stg環境か判断する
func IsStaging() bool {
	return strings.HasSuffix(projectID, "-stg")
}

// IsProd - prod環境か判断する
func IsProd() bool {
	return strings.HasSuffix(projectID, "-prod")
}

// GetProjectID - プロジェクトIDを取得する
func GetProjectID() string {
	id := projectID
	if id == "" {
		return "localProject"
	}
	return id
}

// GetServiceName - サービス名を取得する
func GetServiceName() string {
	if serviceName == "" {
		return "localService"
	}
	return serviceName
}

// GetServiceVersion - バージョンを取得する
func GetServiceVersion() string {
	if serviceVersion == "" {
		return "1.0"
	}
	return serviceVersion
}
