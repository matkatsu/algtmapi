package controller

import (
	"algtmapi/aikatsup"
	"algtmapi/app"

	"github.com/goadesign/goa"
	"github.com/k0kubun/pp"
)

// SearchController implements the Search resource.
type SearchController struct {
	*goa.Controller
}

// NewSearchController creates a Search controller.
func NewSearchController(service *goa.Service) *SearchController {
	return &SearchController{Controller: service.NewController("SearchController")}
}

// GetSearch runs the GetSearch action.
func (c *SearchController) GetSearch(ctx *app.GetSearchSearchContext) error {
	// SearchController_GetSearch: start_implement

	// Put your logic here
	client, err := aikatsup.NewClient("http://aikatsup.com/api/v1/", nil)
	if err != nil {
		return err
	}
	result, err := client.GetSearchResult(ctx.Word)
	if err != nil {
		pp.Println(err)
		return err
	}
	pp.Println(result)
	// TODO: github.com/parnurzeal/gorequest

	res := &app.RomiogakuComAlgtmapiSearch{}
	return ctx.OK(res)
	// SearchController_GetSearch: end_implement
}
