package app

import (
	"context"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/office-aska/lwgc/pkg/perm"
	"github.com/office-aska/lwgc/repository"
)

func authorizer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()

		if strings.HasPrefix(r.URL.Path, "/signin/") {
			return next(c)
		}

		if strings.HasPrefix(r.URL.Path, "/signout/") {
			return next(c)
		}

		if strings.HasPrefix(r.URL.Path, "/queue/") {
			return next(c)
		}

		userID, err := perm.Username(c)
		if err != nil || userID == "" {
			v := url.Values{}
			v.Set("next", r.URL.Path)
			c.Redirect(http.StatusFound, "/signin/?"+v.Encode())
			return nil
		}

		ctx := r.Context()
		fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
		if err != nil {
			return err
		}
		defer fsClient.Close()
		repo := repository.NewUserRepository(fsClient)
		user, err := repo.Get(ctx, userID)
		if err != nil {
			return err
		}
		user.ID = userID
		ctx = context.WithValue(
			ctx, model.User{}, user,
		)
		c.SetRequest(r.WithContext(ctx))
		r = c.Request()
		if user.IsAdmin {
			return next(c)
		}

		if strings.HasPrefix(r.URL.Path, "/organization/") {
			c.Redirect(http.StatusFound, "/signin/")
			return nil
		}

		if r.URL.Path == "/event/" {
			c.Redirect(http.StatusFound, "/signin/")
			return nil
		}

		if r.URL.Path == "/event/create" {
			c.Redirect(http.StatusFound, "/signin/")
			return nil
		}

		ok, err := regexp.MatchString("/event/[0-9a-zA-Z]+/update", r.URL.Path)
		if err != nil {
			return err
		}
		if ok {
			c.Redirect(http.StatusFound, "/signin/")
			return nil
		}

		return next(c)
	}
}
