package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/luxifer/ical"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/teambition/rrule-go"
)

type LineWorksScheduleListReq struct {
	From  time.Time
	To    time.Time
	Token string
}

type LineWorksScheduleListResp struct {
	Result      string              `json:"result"`
	ReturnValue []LineWorksSchedule `json:"returnValue"`
}

type LineWorksSchedule struct {
	UserID     string `json:"userId"`
	CalendarID string `json:"calendarId"`
	Ical       string `json:"ical"`
	ViewURL    string `json:"viewUrl"`
}

type LineWorksEvent struct {
	UID          string
	Source       string
	Location     string
	Created      time.Time
	LastModified time.Time
	StartDate    time.Time
	EndDate      time.Time
	RuleDates    []time.Time
	SummaryDate  string
	Summary      string
	Description  string
	Organizer    string
	ViewURL      string
	Raw          string
}

func SearchLineWorksEvents(listReq *LineWorksScheduleListReq) ([]*LineWorksEvent, error) {
	values := url.Values{}
	values.Add("rangeDateFrom", listReq.From.Format("20060102"))
	values.Add("rangeDateUntil", listReq.To.Format("20060102"))
	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://apis.worksmobile.com/%s/calendar/getScheduleList/V3",
			environ.GetLineWorksAppID()),
		strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("consumerKey", environ.GetLineWorksConsumerKey())
	req.Header.Set("Authorization", "Bearer "+listReq.Token)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonResp := &LineWorksScheduleListResp{}
	err = json.NewDecoder(resp.Body).Decode(jsonResp)
	if err != nil {
		return nil, err
	}

	events := make([]*LineWorksEvent, 0)
	for _, schedule := range jsonResp.ReturnValue {
		calendar, err := ical.Parse(strings.NewReader(schedule.Ical), nil)
		if err != nil {
			// log.Warningf(ctx, "[ical.Parse] error:%v", err)
			continue
		}

		for _, event := range calendar.Events {
			me := &LineWorksEvent{}
			me.ViewURL = schedule.ViewURL
			me.Summary = event.Summary
			me.Description = strings.Replace(event.Description, "\\n", "\n", -1)
			me.StartDate = event.StartDate.In(environ.LocationJST())
			me.EndDate = event.EndDate.In(environ.LocationJST())

			for _, prop := range event.Properties {
				switch prop.Name {
				case "UID":
					me.UID = prop.Value
				case "LOCATION":
					me.Location = prop.Value
				case "CREATED":
					if t, err := time.Parse("20060102T150405Z", prop.Value); err == nil {
						me.Created = t.In(environ.LocationJST())
					}
				case "LAST-MODIFIED":
					if t, err := time.Parse("20060102T150405Z", prop.Value); err == nil {
						me.LastModified = t.In(environ.LocationJST())
					}
				case "RRULE":
					option, err := rrule.StrToROptionInLocation(prop.Value, environ.LocationJST())
					if err != nil {
						return nil, err
					}
					option.Dtstart = me.StartDate
					option.Until = option.Until.In(environ.LocationJST())
					rule, err := rrule.NewRRule(*option)
					if err != nil {
						return nil, err
					}
					me.RuleDates = rule.Between(listReq.From, listReq.To, true)
				case "ORGANIZER":
					for key, param := range prop.Params {
						if key == "CN" {
							for _, v := range param.Values {
								me.Organizer = v
							}
						}
					}
				}
			}

			if len(me.RuleDates) > 0 {
				for _, date := range me.RuleDates {
					if date.Format("20060102") >= listReq.From.Format("20060102") &&
						date.Format("20060102") <= listReq.To.Format("20060102") {
						duration := me.EndDate.Sub(me.StartDate)
						sd := fmt.Sprintf("%sT%s00", date.Format("20060102"), me.StartDate.Format("1504"))
						me.StartDate, _ = time.ParseInLocation("20060102T150405", sd, environ.LocationJST())
						me.EndDate = me.StartDate.Add(duration)
						// log.Debugf(ctx, "    RuleDate:%s from:%s to:%s pass", date.Format("20060102"), from, to)
						break
					}
					// log.Debugf(ctx, "    RuleDate:%s from:%s to:%s skip", date.Format("20060102"), from, to)
				}
			}
			if me.StartDate.Before(listReq.From) && me.EndDate.Before(listReq.From) {
				// log.Debugf(ctx, "    from:%s %s-%s skip", listReq.From, me.StartDate, me.EndDate)
				continue
			}
			if me.StartDate.After(listReq.To) && me.EndDate.After(listReq.To) {
				// log.Debugf(ctx, "    to:%s %s-%s skip", listReq.To, me.StartDate, me.EndDate)
				continue
			}
			if listReq.From.Format("20060102") == me.StartDate.Format("20060102") &&
				listReq.From.Format("20060102") == me.EndDate.Format("20060102") {
				me.SummaryDate = fmt.Sprintf("%s〜%s",
					me.StartDate.Format("15:04"),
					me.EndDate.Format("15:04"))
			} else {
				me.SummaryDate = fmt.Sprintf("%s〜\n%s",
					me.StartDate.Format("2006/01/02 15:04"),
					me.EndDate.Format("2006/01/02 15:04"))
			}

			events = append(events, me)
			break
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].StartDate.Before(events[j].StartDate)
	})

	return events, nil
}
