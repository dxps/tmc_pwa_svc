package comps

import (
	_ "embed"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Logo struct {
	app.Compo
}

func (l *Logo) Render() app.UI {
	return app.Div().Class("w-[24px] h-[20px]").Body(
		app.Raw(LOGO_ICON),
	)
}
