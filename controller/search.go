package controller

import (
	"algtmapi/aikatsup"
	"algtmapi/app"
	"algtmapi/env"

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
	client := aikatsup.Client{
		BaseURL: env.Aikatup.BaseURL,
	}
	result, err := client.GetSearchResult(ctx.Word)

	if err != nil {
		goa.NewErrorClass(err.Error(), 400)
	}

	var list []*app.SearchItemType
	for _, v := range result {
		tmp := &app.SearchItemType{
			ID:   v.ID,
			URL:  v.Image.URL,
			Word: v.Word,
		}
		list = append(list, tmp)
	}
	res := &app.RomiogakuComAlgtmapiSearch{Result: list}
	return ctx.OK(res)
}
