package googlecalendar

import (
	"crypto/subtle"
	"errors"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/domain/service"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/office-aska/lwgc/repository"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
)

// Root /googlecalendar/
func Root(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	csrf, err := r.Cookie("_csrf")
	if err != nil {
		return err
	}
	conf := service.GetGoogleCalendarConfig(ctx, scheme()+"://"+r.Host+"/googlecalendar/callback")

	// ApprovalForce ないと refresh token くれない
	// refs. https://godoc.org/golang.org/x/oauth2#AuthCodeOption
	authURL := conf.AuthCodeURL(csrf.Value, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return c.Redirect(http.StatusFound, authURL)
}

// Callback /googlecalendar/callback
func Callback(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	csrf, err := r.Cookie("_csrf")
	if err != nil {
		return err
	}

	user, ok := ctx.Value(model.User{}).(*model.User)
	if !ok {
		return errors.New("Missing sign in user")
	}

	code := r.FormValue("code")
	state := r.FormValue("state")
	if subtle.ConstantTimeCompare([]byte(state), []byte(csrf.Value)) != 1 {
		return c.Redirect(http.StatusFound, "/?google_calendar_error=invalid_statet")
	}
	conf := service.GetGoogleCalendarConfig(ctx, scheme()+"://"+r.Host+"/googlecalendar/callback")
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return xerrors.Errorf("OAuth Error:%w", err)
	}

	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	repo := repository.NewUserRepository(fsClient)
	param := &repository.UserUpdateParam{
		GoogleCalendarAccessToken:  tok.AccessToken,
		GoogleCalendarRefreshToken: tok.RefreshToken,
		GoogleCalendarTokenType:    tok.TokenType,
		GoogleCalendarExpiry:       tok.Expiry,
	}
	err = repo.StrictUpdate(ctx, user.ID, param)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/")
}

// Post /googlecalendar/
func Post(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()

	id := r.FormValue("id")
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
		GoogleCalendarEmail: id,
	}
	err = repo.StrictUpdate(ctx, user.ID, param)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/")
}

func scheme() string {
	if environ.IsLocal() {
		return "http"
	}
	return "https"
}
