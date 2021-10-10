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

func UserRoot(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	oid := c.Param("oid")
	orep := repository.NewOrganizationRepository(fsClient)
	org, err := orep.Get(ctx, oid)
	if err != nil {
		return xerrors.Errorf("OrganizationRepository#Get error:%w", err)
	}

	param := &repository.OrganizationUserSearchParam{
		CursorLimit: 20,
	}
	parent := orep.GetDocRef(oid)
	rows, err := repository.NewOrganizationUserRepository(fsClient, parent).Search(ctx, param, nil)
	if err != nil {
		return xerrors.Errorf("OrganizationUserRepository#Search error:%w", err)
	}

	data := map[string]interface{}{}
	data["org"] = org
	data["rows"] = rows
	return c.Render(http.StatusOK, templateDir+"/user/index.html", data)
}

func UserView(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	oid := c.Param("oid")
	orep := repository.NewOrganizationRepository(fsClient)
	org, err := orep.Get(ctx, oid)
	if err != nil {
		return xerrors.Errorf("OrganizationRepository#Get error:%w", err)
	}

	parent := orep.GetDocRef(oid)
	uid := c.Param("uid")
	row, err := repository.NewOrganizationUserRepository(fsClient, parent).Get(ctx, uid)
	if err != nil {
		return xerrors.Errorf("firestore.OrganizationUserRetrieveByID error:%w", err)
	}

	data := map[string]interface{}{}
	data["org"] = org
	data["row"] = row
	return c.Render(http.StatusOK, templateDir+"/user/view.html", data)
}

func UserCreate(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	oid := c.Param("oid")
	orep := repository.NewOrganizationRepository(fsClient)
	org, err := orep.Get(ctx, oid)
	if err != nil {
		return xerrors.Errorf("OrganizationRepository#Get error:%w", err)
	}

	data := map[string]interface{}{}
	data["org"] = org
	data["row"] = &model.OrganizationUser{}
	return c.Render(http.StatusOK, templateDir+"/user/register.html", data)
}

func UserCreatePost(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	oid := c.Param("oid")
	orep := repository.NewOrganizationRepository(fsClient)

	req := &model.OrganizationUser{
		PrimaryEmail: r.FormValue("PrimaryEmail"),
	}
	parent := orep.GetDocRef(oid)
	uid, err := repository.NewOrganizationUserRepository(fsClient, parent).Insert(ctx, req)
	if err != nil {
		return xerrors.Errorf("firestore.OrganizationUserCreate error:%w", err)
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/organization/%s/user/%s/", oid, uid))
}

func UserUpdate(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	oid := c.Param("oid")
	orep := repository.NewOrganizationRepository(fsClient)
	org, err := orep.Get(ctx, oid)
	if err != nil {
		return xerrors.Errorf("OrganizationRepository#Get error:%w", err)
	}

	parent := orep.GetDocRef(oid)

	uid := c.Param("uid")
	row, err := repository.NewOrganizationUserRepository(fsClient, parent).Get(ctx, uid)
	if err != nil {
		return xerrors.Errorf("firestore.OrganizationUserRetrieveByID error:%w", err)
	}

	data := map[string]interface{}{}
	data["org"] = org
	data["row"] = row
	return c.Render(http.StatusOK, templateDir+"/user/register.html", data)
}

func UserUpdatePost(c echo.Context) error {
	r := c.Request()
	ctx := r.Context()
	fsClient, err := firestore.NewClient(ctx, environ.GetProjectID())
	if err != nil {
		return xerrors.Errorf("firestore.NewClient error:%w", err)
	}
	defer fsClient.Close()

	oid := c.Param("oid")
	orep := repository.NewOrganizationRepository(fsClient)

	parent := orep.GetDocRef(oid)

	uid := c.Param("uid")
	param := &repository.OrganizationUserUpdateParam{
		PrimaryEmail: r.FormValue("PrimaryEmail"),
	}
	err = repository.NewOrganizationUserRepository(fsClient, parent).StrictUpdate(ctx, uid, param)
	if err != nil {
		return xerrors.Errorf("OrganizationUserRepository#Update error:%w", err)
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/organization/%s/user/%s/", oid, uid))
}
