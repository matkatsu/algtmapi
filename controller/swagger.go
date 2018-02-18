package controller

import (
	"algtmapi/swagger/swaggerassets"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/goadesign/goa"

	"github.com/mattetti/filebuffer"
	"golang.org/x/net/context"
)

// SwaggerController implements the swagger resource.
type SwaggerController struct {
	*goa.Controller
}

// NewSwaggerController creates a swagger controller.
func NewSwaggerController(service *goa.Service) *SwaggerController {
	return &SwaggerController{Controller: service.NewController("SwaggerController")}
}

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl *SwaggerController) {
	initService(service)
	var h goa.Handler

	h = GoaGobinAssetHandler("/swagger.json", "swagger/specs/swagger.json")
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/specs/swagger.json", "route", "GET /swagger.json")

	h = GoaGobinAssetHandler("/swaggerui/*filepath", "swagger/swaggerui/dist")
	service.Mux.Handle("GET", "/swaggerui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swaggerui/dist", "route", "GET /swaggerui/*filepath")

	h = GoaGobinAssetHandler("/swaggerui/", "swagger/swaggerui/dist/index.html")
	service.Mux.Handle("GET", "/swaggerui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swaggerui/dist/index.html", "route", "GET /swaggerui/")
}

// GoaGobinAssetHandler goa.Handler for process assets
func GoaGobinAssetHandler(path, filename string) goa.Handler {
	var wc string
	if idx := strings.LastIndex(path, "/*"); idx > -1 && idx < len(path)-1 {
		wc = path[idx+2:]
		if strings.Contains(wc, "/") {
			wc = ""
		}
	}
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		fname := filename
		if len(wc) > 0 {
			if m, ok := goa.ContextRequest(ctx).Params[wc]; ok {
				fname = filepath.Join(filename, m[0])
			}
		}
		goa.LogInfo(ctx, "serve file", "name", fname, "route", req.URL.Path)

		data, err := swaggerassets.Asset(fname)
		if err != nil {
			return goa.ErrInvalidFile(err)
		}
		f := filebuffer.New(data)

		d, err := swaggerassets.AssetInfo(fname)
		if err != nil {
			return goa.ErrInvalidFile(err)
		}
		http.ServeContent(rw, req, d.Name(), d.ModTime(), f)
		return nil
	}
}
