package api4

import (
	"net/http"

	"mattermost-server/model"
	"mattermost-server/utils"
)

func (api *API) InitOrgs() {
	api.BaseRoutes.Orgs.Handle("/createcompany", api.ApiHandler(createCompany)).Methods("POST")
	api.BaseRoutes.Orgs.Handle("/getcompany", api.ApiHandler(getCompany)).Methods("POST")
	api.BaseRoutes.Orgs.Handle("/createorg", api.ApiHandler(createOrg)).Methods("POST")
	api.BaseRoutes.Orgs.Handle("/getorg", api.ApiHandler(getOrg)).Methods("POST")
	api.BaseRoutes.Orgs.Handle("/delorg", api.ApiHandler(delOrg)).Methods("POST")
	api.BaseRoutes.Orgs.Handle("/updateorg", api.ApiHandler(updateOrg)).Methods("POST")
}

func createCompany(c *Context, w http.ResponseWriter, r *http.Request) {
	company := model.MapFromJson(r.Body)
	if company == nil {
		c.SetInvalidParam("company")
		return
	}
	org := &model.Org{
		Key:  company["key"],
		Name: company["name"],
	}
	if utils.InValidString(org.Name) {
		c.Err = model.NewAppError("createCompany", "api.orgs.createCompany.invalidstring.app_error", nil, "string is invalid", http.StatusForbidden)
		return
	}
	// if !c.App.SessionHasPermissionTo(c.Session, model.PERMISSION_CREATE_COMPANY) {
	// 	c.Err = model.NewAppError("createCompany", "api.company.is_company_creation_allowed.disabled.app_error", nil, "", http.StatusForbidden)
	// 	return
	// }
	result, err := c.App.CreateCompany(org)
	if err != nil {
		c.Err = err
		return
	}

	// Don't sanitize the company here since the user will be a company admin and their session won't reflect that yet

	// w.WriteHeader(http.StatusCreated)
	w.Write([]byte(model.MapToJson(result)))
}

func getCompany(c *Context, w http.ResponseWriter, r *http.Request) {
	props := model.MapFromJson(r.Body)
	key := props["key"]
	if key == "" {
		c.Err = model.NewAppError("getCompany", "api.orgs.getCompany.invalidstring.app_error", nil, "string is nil", http.StatusForbidden)
		return
	}
	result, err := c.App.GetCompany(key)
	if err != nil {
		c.Err = err
		return
	}
	w.Write([]byte(model.MapToJson(result)))
}

func createOrg(c *Context, w http.ResponseWriter, r *http.Request) {
	orgPr := model.MapFromJson(r.Body)
	if orgPr == nil {
		c.SetInvalidParam("orgPr")
		return
	}
	parent, name := orgPr["parent"], orgPr["name"]
	if utils.InValidString(name) {

		c.Err = model.NewAppError("createOrg", "api.orgs.createOrg.invalidstring.app_error", nil, "string is invalid", http.StatusForbidden)
		return
	}
	result, err := c.App.CreateOrg(parent, name)
	if err != nil {
		c.Err = err
		return
	}
	w.Write([]byte(model.MapToJson(result)))
}

func getOrg(c *Context, w http.ResponseWriter, r *http.Request) {
	orgPr := model.MapFromJson(r.Body)
	if orgPr == nil {
		c.SetInvalidParam("orgPr")
		return
	}
	key := orgPr["key"]
	if key == "" {
		c.Err = model.NewAppError("getOrg", "api.orgs.getOrg.invalidstring.app_error", nil, "string is nil", http.StatusForbidden)
	}
	result, err := c.App.GetOrgNode(key)
	if err != nil {
		c.Err = err
		return
	}
	w.Write([]byte(model.MapToJson(result)))
}
func delOrg(c *Context, w http.ResponseWriter, r *http.Request) {
	domain := model.MapFromJson(r.Body)
	if domain["key"] == "" {
		c.Err = model.NewAppError("delOrg", "api.orgs.delOrg.invalidstring.app_error", nil, "string is nil", http.StatusForbidden)
	}
	result, err := c.App.DelOrg(domain["key"])
	if err != nil {
		c.Err = err
		return
	}
	w.Write([]byte(model.MapToJson(result)))
}

func updateOrg(c *Context, w http.ResponseWriter, r *http.Request) {
	domain := model.MapFromJson(r.Body)
	if domain["key"] == "" || domain["name"] == "" {
		c.Err = model.NewAppError("updateOrg", "api.orgs.UpdateOrg.invalidstring.app_error", nil, "string is nil", http.StatusForbidden)
		return
	}
	if utils.InValidString(domain["name"]) {
		c.Err = model.NewAppError("updateOrg", "api.orgs.UpdateOrg.invalidstring.app_error", nil, "string is invalid", http.StatusForbidden)
		return
	}
	result, err := c.App.UpdateOrg(domain["key"], domain["name"])
	if err != nil {
		c.Err = err
		return
	}
	w.Write([]byte(model.MapToJson(result)))
}
