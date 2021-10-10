package perm

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/office-aska/lwgc/repository"
	"golang.org/x/xerrors"
)

var cookieName = "SID"

// SignIn ...
func SignIn(c echo.Context, userID string) error {
	ctx := c.Request().Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return err
	}
	defer fsClient.Close()
	sub := &model.Session{
		UserID: userID,
	}
	repo := repository.NewSessionRepository(fsClient)
	id, err := repo.Insert(ctx, sub)
	if err != nil {
		return xerrors.Errorf("firestore.SessionCreate error:%w", err)
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    id,
		Path:     "/",
		Secure:   !environ.IsLocal(),
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return nil
}

// SignOut ...
func SignOut(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     cookieName,
		Path:     "/",
		Secure:   !environ.IsLocal(),
		HttpOnly: true,
		MaxAge:   -1,
	})

	ctx := c.Request().Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return err
	}
	defer fsClient.Close()

	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return err
	}

	sid := cookie.Value
	repo := repository.NewSessionRepository(fsClient)
	_ = repo.DeleteByID(ctx, sid)

	return nil
}

// Username ...
func Username(c echo.Context) (string, error) {
	ctx := c.Request().Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return "", err
	}
	defer fsClient.Close()

	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return "", err
	}

	sid := cookie.Value
	repo := repository.NewSessionRepository(fsClient)
	item, err := repo.Get(ctx, sid)
	if err != nil {
		return "", err
	}

	return item.UserID, err
}
