package comps

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Navbar struct {
	app.Compo
}

func (n *Navbar) Render() app.UI {
	return app.Nav().
		Class("absolute w-full px-4 py-2 flex justify-between items-center bg-white z-40").
		Body(
			app.A().
				Class("text-3xl font-bold leading-none").
				Href("/").
				Body(
					&Logo{},
				),
		)
}
