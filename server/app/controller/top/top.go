package top

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/domain/service"
	"github.com/office-aska/lwgc/pkg/environ"
	"google.golang.org/api/calendar/v3"
)

// Root トップページ
func Root(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	data := map[string]interface{}{}

	user, ok := ctx.Value(model.User{}).(*model.User)
	if ok {

		if user.GoogleCalendarAccessToken != "" {
			conf := service.GetGoogleCalendarConfig(ctx, scheme()+"://"+r.Host+"/googlecalendar/callback")
			client := conf.Client(ctx, user.GetGoogleCalendarToken())
			srv, err := calendar.New(client)
			if err != nil {
				return err
			}
			list, err := srv.CalendarList.List().Do()
			if err != nil {
				return err
			}
			data["calendarListItems"] = list.Items
		}

		type Row struct {
			Date                 time.Time
			LineWorksEvents      []*service.LineWorksEvent
			GoogleCalendarEvents []*calendar.Event
		}
		rows := []*Row{}
		now := environ.NowJST()
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, environ.LocationJST())
		for i := 0; i < 3; i++ {
			row := &Row{
				Date: start.AddDate(0, 0, i),
			}
			if user.LINEWorksAccessToken != "" {
				from := start.AddDate(0, 0, i)
				to := start.AddDate(0, 0, i+1)
				req := &service.LineWorksScheduleListReq{
					From:  from,
					To:    to,
					Token: user.LINEWorksAccessToken,
				}
				events, err := service.SearchLineWorksEvents(req)
				if err != nil {
					return err
				}
				row.LineWorksEvents = events
			}

			if user.GoogleCalendarEmail != "" {
				conf := service.GetGoogleCalendarConfig(ctx, scheme()+"://"+r.Host+"/googlecalendar/callback")
				client := conf.Client(ctx, user.GetGoogleCalendarToken())
				srv, err := calendar.New(client)
				if err != nil {
					return err
				}
				from := start.AddDate(0, 0, i)
				to := start.AddDate(0, 0, i+1)
				calendarEvents, err := srv.Events.List(user.GoogleCalendarEmail).
					SingleEvents(true).
					TimeMin(from.Format(time.RFC3339)).
					TimeMax(to.Format(time.RFC3339)).
					OrderBy("starttime").
					Do()
				if err != nil {
					return err
				}

				fmt.Printf("%+v", calendarEvents.Items)
				row.GoogleCalendarEvents = calendarEvents.Items
			}

			rows = append(rows, row)
		}
		data["rows"] = rows
	}

	return c.Render(http.StatusOK, "index.html", data)
}

func scheme() string {
	if environ.IsLocal() {
		return "http"
	}
	return "https"
}
