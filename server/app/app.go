package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/office-aska/lwgc/server/app/controller/googlecalendar"
	"github.com/office-aska/lwgc/server/app/controller/lineworks"
	"github.com/office-aska/lwgc/server/app/controller/signin"
	"github.com/office-aska/lwgc/server/app/controller/top"
)

func Bootstrap() {

	e := echo.New()
	http.Handle("/", e)
	e.Renderer = NewTemplate()

	e.Use(authorizer)
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		Skipper: func(c echo.Context) bool {
			// if strings.HasPrefix(c.Request().URL.Path, "/tq/") {
			// 	return true
			// }
			// return c.Request().URL.Path == "/media/upload/callback"
			return false
		},
		TokenLookup:    "form:csrf_token",
		CookiePath:     "/",
		CookieMaxAge:   365 * 24 * 60 * 60,
		CookieSecure:   !environ.IsLocal(),
		CookieHTTPOnly: true,
	}))
	SetErrorHandler(e)

	// memorystore.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}
	// log.Printf("Listening on port %s proj:%s isLocal:%v", port, env.ProjectID, env.IsLocal())

	e.Static("/static", "app/static")

	e.GET("/", top.Root)
	e.GET("/signin/", signin.Root)
	e.GET("/signin/callback", signin.Callback)
	e.GET("/signout/", signin.SignOut)
	e.GET("/lineworks/", lineworks.Root)
	e.GET("/lineworks/callback", lineworks.Callback)
	e.GET("/googlecalendar/", googlecalendar.Root)
	e.POST("/googlecalendar/", googlecalendar.Post)
	e.GET("/googlecalendar/callback", googlecalendar.Callback)
	// e.GET("/organization/", organization.Root)
	// e.GET("/organization/create", organization.Create)
	// e.POST("/organization/create", organization.CreatePost)
	// e.GET("/organization/:id/", organization.View)
	// e.GET("/organization/:id/update", organization.Update)
	// e.POST("/organization/:id/update", organization.UpdatePost)
	// e.GET("/organization/:oid/user/", organization.UserRoot)
	// e.GET("/organization/:oid/user/create", organization.UserCreate)
	// e.POST("/organization/:oid/user/create", organization.UserCreatePost)
	// e.GET("/organization/:oid/user/:uid/", organization.UserView)
	// e.GET("/organization/:oid/user/:uid/update", organization.UserUpdate)
	// e.POST("/organization/:oid/user/:uid/update", organization.UserUpdatePost)
	// e.GET("/event/", event.Root)
	// e.GET("/event/create", event.Create)
	// e.POST("/event/create", event.CreatePost)
	// e.GET("/event/:id/", event.View)
	// e.GET("/event/:id/update", event.Update)
	// e.POST("/event/:id/update", event.UpdatePost)
	// e.GET("/event/:id/notify/", event.NotifyRoot)
	// e.GET("/event/:id/notify/create", event.NotifyCreate)
	// e.POST("/event/:id/notify/create", event.NotifyCreatePost)
	// e.GET("/event/:id/notify/:nid/", event.NotifyView)
	// e.GET("/event/:id/notify/:nid/update", event.NotifyUpdate)
	// e.POST("/event/:id/notify/:nid/update", event.NotifyUpdatePost)
	// e.POST("/event/:id/notify/:nid/delete", event.NotifyDeletePost)
	// e.POST("/queue/event/:id/notify/:nid/send", event.NotifySend)
	// e.GET("/event/:id/booth/", event.BoothRoot)
	// e.GET("/event/:id/booth/create", event.BoothCreate)
	// e.POST("/event/:id/booth/create", event.BoothCreatePost)
	// e.GET("/event/:id/booth/:bid/", event.BoothView)
	// e.GET("/event/:id/booth/:bid/update", event.BoothUpdate)
	// e.POST("/event/:id/booth/:bid/update", event.BoothUpdatePost)
	// e.GET("/event/:id/booth/:bid/request/", event.BoothRequestRoot)
	// e.POST("/event/:id/booth/:bid/request/generate", event.BoothRequestGenerate)
	// e.POST("/event/:id/booth/:bid/request/auto", event.BoothRequestInstantGenerate)
	// e.GET("/event/:id/booth/:bid/request/qrscan", event.BoothRequestQRScan)
	// e.POST("/event/:id/booth/:bid/request/qrscan", event.BoothRequestQRScanPost)
	// e.POST("/queue/event/:id/booth/:bid/call", event.CallNotifySend)
	// e.GET("/event/:id/booth/:bid/request/:rid/", event.BoothRequestView)
	// e.POST("/event/:id/booth/:bid/request/:rid/", event.BoothRequestPost)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
