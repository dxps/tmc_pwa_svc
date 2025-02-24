package server

import (
	shttp "github.com/dxps/tmc-pwa/internal/shared/http"
	"github.com/dxps/tmc-pwa/internal/ui/pages"
	"github.com/dxps/tmc-pwa/internal/ui/pages/attr_def"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func initRoutes(apiClient *shttp.ApiClient) {

	app.Route("/", func() app.Composer { return &pages.Homepage{} })
	app.Route("/about", func() app.Composer { return pages.NewAboutPage(apiClient) })
	app.Route("/admin", func() app.Composer { return pages.NewAdminPage() })
	app.Route("/definitions/attributes", func() app.Composer { return attr_def.NewAttributeDefListPage(apiClient) })
}
