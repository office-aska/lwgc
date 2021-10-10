package organization

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/office-aska/lwgc/domain/model"
	"github.com/office-aska/lwgc/pkg/environ"
	"github.com/office-aska/lwgc/repository"
	"golang.org/x/xerrors"
)

const templateDir = "organization"

func Root(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	repo := repository.NewOrganizationRepository(fsClient)
	param := &repository.OrganizationSearchParam{
		CursorLimit: 20,
		// CreatedAt:   repository.NewQueryChainer().Desc(),
	}
	rows, err := repo.Search(ctx, param, nil)
	if err != nil {
		return xerrors.Errorf("OrganizationRepository#Search error:%w", err)
	}

	data := map[string]interface{}{}
	data["rows"] = rows
	return c.Render(http.StatusOK, templateDir+"/index.html", data)
}

func View(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	id := c.Param("id")
	repo := repository.NewOrganizationRepository(fsClient)
	row, err := repo.Get(ctx, id)
	if err != nil {
		return xerrors.Errorf("OrganizationRepository#Get error:%w", err)
	}

	data := map[string]interface{}{}
	data["row"] = row
	return c.Render(http.StatusOK, templateDir+"/view.html", data)
}

func Create(c echo.Context) error {
	data := map[string]interface{}{}

	data["row"] = &model.Organization{}
	return c.Render(http.StatusOK, templateDir+"/register.html", data)
}

func CreatePost(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	req := &model.Organization{
		Name: r.FormValue("Name"),
	}
	repo := repository.NewOrganizationRepository(fsClient)
	id, err := repo.Insert(ctx, req)
	if err != nil {
		return xerrors.Errorf("firestore.OrganizationCreate error:%w", err)
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/organization/%s/", id))
}

func Update(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	id := c.Param("id")
	repo := repository.NewOrganizationRepository(fsClient)
	row, err := repo.Get(ctx, id)
	if err != nil {
		return xerrors.Errorf("OrganizationRepository#Get error:%w", err)
	}

	data := map[string]interface{}{}
	data["row"] = row
	return c.Render(http.StatusOK, templateDir+"/register.html", data)
}

func UpdatePost(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	id := c.Param("id")
	name := r.FormValue("Name")
	repo := repository.NewOrganizationRepository(fsClient)
	param := &repository.OrganizationUpdateParam{
		Name: name,
	}
	err = repo.StrictUpdate(ctx, id, param)
	if err != nil {
		return xerrors.Errorf("OrganizationRepository#StrictUpdate error:%w", err)
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/organization/%s/", id))
}
