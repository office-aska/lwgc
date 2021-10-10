package service

import (
	"context"

	"github.com/office-aska/lwgc/pkg/environ"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func GetGoogleCalendarConfig(ctx context.Context, callbackURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     environ.GetGoogleOAuthClientID(),
		ClientSecret: environ.GetGoogleOAuthClientSecret(),
		Scopes:       []string{calendar.CalendarScope},
		Endpoint:     google.Endpoint,
		RedirectURL:  callbackURL,
	}
}
