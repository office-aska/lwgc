package top

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/pkg/extract"
)

// Root トップページ
func Root(c echo.Context) error {
	data := map[string]interface{}{}

	// fsClient, err := firestore.NewClient(c.Request().Context())
	// if err != nil {
	// 	return xerrors.Errorf("firestore.NewClient error:%w", err)
	// }
	// defer fsClient.Close()
	// log.Infof("firestore.Newclient %v", fsClient)

	userID, ok := extract.UserID(c.Request().Context())
	if ok {
		log.Printf("userID: %s", userID)
	}

	return c.Render(http.StatusOK, "index.html", data)
}
