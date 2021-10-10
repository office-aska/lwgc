package signin

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
	"regexp"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/office-aska/lwgc/pkg/perm"
	"github.com/office-aska/lwgc/repository"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

func getConfig(r *http.Request) *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: scheme() + "://" + r.Host + "/signin/callback",
	}
	return config
}

var nextCookieName = "signin-state-next"
var googleCookieName = "signin-state-google"
var googleScope = "User.Read"

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func scheme() string {
	if environ.IsLocal() {
		return "http"
	}
	return "https"
}

// GoogleTokenInfo GoogleTokenInfo
type GoogleTokenInfo struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	RefreshToken     string `json:"refresh_token"`
	IDToken          string `json:"id_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	Timestamp        string `json:"timestamp"`
	TraceID          string `json:"trace_id"`
	CorrelationID    string `json:"correlation_id"`
}

// GoogleUser ...
type GoogleUser struct {
	OdataContext      string   `json:"@odata.context"`
	BusinessPhones    []string `json:"businessPhones"`
	DisplayName       string   `json:"displayName"`
	GivenName         string   `json:"givenName"`
	ID                string   `json:"id"`
	JobTitle          string   `json:"jobTitle"`
	Mail              string   `json:"mail"`
	MobilePhone       string   `json:"mobilePhone"`
	OfficeLocation    string   `json:"officeLocation"`
	PreferredLanguage string   `json:"preferredLanguage"`
	Surname           string   `json:"surname"`
	UserPrincipalName string   `json:"userPrincipalName"`
}

// Root ...
func Root(c echo.Context) error {
	r := c.Request()

	b, err := generateRandomBytes(32)
	if err != nil {
		return err
	}
	state := hex.EncodeToString(b)

	c.SetCookie(&http.Cookie{
		Name:     googleCookieName,
		Value:    state,
		Path:     "/signin/",
		Secure:   !environ.IsLocal(),
		HttpOnly: true,
	})

	c.SetCookie(&http.Cookie{
		Name:     nextCookieName,
		Value:    r.FormValue("next"),
		Path:     "/signin/",
		Secure:   !environ.IsLocal(),
		HttpOnly: true,
	})

	url := getConfig(r).AuthCodeURL(state, oauth2.ApprovalForce)
	return c.Redirect(http.StatusFound, url)
}

// Callback ...
func Callback(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()

	code := r.FormValue("code")
	state := r.FormValue("state")

	stateCookie, err := r.Cookie(googleCookieName)
	if err != nil {
		return err
	}

	nextCookie, err := r.Cookie(nextCookieName)
	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare([]byte(state), []byte(stateCookie.Value)) != 1 {
		return errors.New("Invalid state")
	}

	config := getConfig(r)
	tok, err := config.Exchange(ctx, code)
	if err != nil {
		return err
	}

	if tok.Valid() == false {
		return errors.New("Invalid token")
	}

	service, err := v2.New(config.Client(ctx, tok))
	if err != nil {
		return err
	}

	tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	if err != nil {
		return err
	}

	email := tokenInfo.Email
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return err
	}
	defer fsClient.Close()
	repo := repository.NewUserRepository(fsClient)
	param := &repository.UserSearchParam{
		PrimaryEmail: repository.NewQueryChainer().Equal(email),
	}
	rows, err := repo.Search(ctx, param, nil)
	if err != nil {
		return err
	}
	if len(rows) == 0 && email == "aska@nl-plus.co.jp" {
		sub := &model.User{
			Name:         "aska",
			PrimaryEmail: "aska@nl-plus.co.jp",
			IsAdmin:      true,
		}
		id, err := repo.Insert(ctx, sub)
		if err != nil {
			return err
		}
		err = perm.SignIn(c, id)
		if err != nil {
			return err
		}
	} else if len(rows) == 1 {
		err = perm.SignIn(c, rows[0].ID)
		if err != nil {
			return err
		}
	} else {
		return c.Redirect(http.StatusFound, "/")
	}

	if next := nextCookie.Value; next != "" {
		ok, err := regexp.MatchString("/event/[0-9a-zA-Z]+/", next)
		if err != nil {
			return err
		}
		if ok {
			return c.Redirect(http.StatusFound, next)
		}
	}

	return c.Redirect(http.StatusFound, "/")
}

// SignOut /signout/
func SignOut(c echo.Context) error {
	if err := perm.SignOut(c); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/signin/")
}
