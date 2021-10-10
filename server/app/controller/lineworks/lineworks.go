package lineworks

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/office-aska/lwgc/repository"
	"golang.org/x/xerrors"
)

// Root generate token
func Root(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()

	b, err := generateRandomBytes(32)
	if err != nil {
		return err
	}
	state := hex.EncodeToString(b)

	user, ok := ctx.Value(model.User{}).(*model.User)
	if !ok {
		return errors.New("Missing sign in user")
	}

	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	repo := repository.NewUserRepository(fsClient)
	param := &repository.UserUpdateParam{
		LINEWorksState: state,
	}
	err = repo.StrictUpdate(ctx, user.ID, param)
	if err != nil {
		return err
	}

	q := url.Values{}
	q.Set("client_id", environ.GetLineWorksConsumerKey())
	q.Set("redirect_uri", scheme()+"://"+r.Host+r.RequestURI+"callback")
	q.Set("state", state)
	q.Set("domain", user.Domain)
	redirectURL := fmt.Sprintf("https://auth.worksmobile.com/ba/%s/service/authorize?%s", environ.GetLineWorksAppID(), q.Encode())
	return c.Redirect(http.StatusFound, redirectURL)
}

// Callback callback page
func Callback(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()

	user, ok := ctx.Value(model.User{}).(*model.User)
	if !ok {
		return errors.New("Missing sign in user")
	}

	code := r.FormValue("code")
	state := r.FormValue("state")
	if user.LINEWorksState != state {
		return errors.New("invalid state")
	}

	values := url.Values{}
	values.Add("client_id", environ.GetLineWorksConsumerKey())
	values.Add("domain", user.Domain)
	values.Add("code", code)
	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://auth.worksmobile.com/ba/%s/service/token",
			environ.GetLineWorksAppID()),
		strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	type token struct {
		ErrorCode    string `json:"errorCode"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpireIn     string `json:"expire_in"`
	}
	t := &token{}
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}
	if t.AccessToken == "" {
		return fmt.Errorf("Error: %+v", t)
	}

	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	repo := repository.NewUserRepository(fsClient)
	param := &repository.UserUpdateParam{
		LINEWorksState:        "",
		LINEWorksAccessToken:  t.AccessToken,
		LINEWorksRefreshToken: t.RefreshToken,
	}
	err = repo.StrictUpdate(ctx, user.ID, param)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/")
}

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
