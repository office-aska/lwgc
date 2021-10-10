package ctxkeys

type contextKey = string

// JWTClaims - echo context key for middleware.JWTConfig
const JWTClaims string = "ctxkeys_jwt_claims"

// SessionID - echo context key for claims.jti
const SessionID string = "ctxkeys_session_id"

// UserID - echo context key for claims.uid
const UserID string = "ctxkeys_user_id"

// ExternalUserID - echo context key for claims.eui
const ExternalUserID string = "ctxkeys_external_user_id"

// Role - echo context key for claims.role
const Role string = "ctxkeys_role"

// Issuer - echo context key for claims.iss
const Issuer string = "ctxkeys_issuer"

// UserLanguage - echo context key for i18n.
const UserLanguage string = "ctxkeys_user_language"

// LoggerKey - log.Logrus
const LoggerKey contextKey = "ctxkeys_logger"

// ValidatorKey - *validator.Validate
const ValidatorKey contextKey = "ctxkeys_validator"

// RequestBodyKey - []byte
const RequestBodyKey contextKey = "ctxkeys_request_body"

// ResponseBodyKey - []byte
const ResponseBodyKey contextKey = "ctxkeys_response_body"

type (
	// SessionIDKey - context key for SessionID
	SessionIDKey struct{}
	// UserIDKey - context key for UserID
	UserIDKey struct{}
	// ExternalUserIDKey - context key for ExternalUserID
	ExternalUserIDKey struct{}
	// RoleKey - context key for Role
	RoleKey struct{}
	// IssuerKey - context key for JWT Issuer
	IssuerKey struct{}
)
