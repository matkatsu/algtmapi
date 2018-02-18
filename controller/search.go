package controller

import (
	"algtmapi/app"

	"github.com/goadesign/goa"
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

	res := &app.RomiogakuComAlgtmapiSearch{}
	return ctx.OK(res)
	// SearchController_GetSearch: end_implement
}
