package comps

import "github.com/maxence-charriere/go-app/v10/pkg/app"

const (
	navbarLinkCss = "text-sm text-gray-600 py-2 px-4 hover:bg-gray-100 rounded-lg transition duration-200"
)

type Navbar struct {
	app.Compo
}

func (n *Navbar) Render() app.UI {
	return app.Nav().
		Class("absolute w-full px-4 py-1 flex justify-between items-center bg-white z-40").
		Body(
			app.A().
				Class("text-gray-500 hover:text-gray-500 py-2").
				Href("/").
				Body(&Logo{}),
			app.Ul().
				Class(`hidden absolute top-1/2 sm:left-1/3 sm:pl-16 md:left-1/2 lg:left-1/2
                    transform -translate-y-1/2 -translate-x-1/2
                    sm:flex sm:mx-auto sm:flex sm:items-center sm:w-auto sm:space-x-3 lg:space-x-6`).Body(
				app.Li().Body(
					app.A().Href("/").Text("Home").Class(navbarLinkCss),
				),
				&NavSep{},
				app.Li().Body(
					app.A().Href("/data-mgmt").Text("Data Management").Class(navbarLinkCss),
				),
			),
			&NavUserMenu{},
		)
}

type NavSep struct {
	app.Compo
}

func (n *NavSep) Render() app.UI {
	return app.Li().
		Class("text-gray-300").
		Body(
			app.Raw(`
				<svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" 
				    class="w-4 h-4 current-fill" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                          d="M12 5v0m0 7v0m0 7v0m0-13a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" 
					/>
                </svg>
			`),
		)
}
