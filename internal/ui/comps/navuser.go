package comps

import (
	"fmt"
	"log/slog"

	"github.com/dxps/tmc-pwa/internal/ui/logic/auth"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const (
	DropdownLinkCss   = "flex items-center text-[#333] hover:bg-gray-100 hover:text-orange-600 text-sm cursor-pointer"
	ShowDropdownState = "showDropdown"
)

type NavUserMenu struct {
	app.Compo
	inited       bool
	token        *string
	showDropdown bool
}

func (n *NavUserMenu) OnMount(ctx app.Context) {

	if ctx.LocalStorage().Contains(auth.LSK_USER_TOKEN) {
		if err := ctx.LocalStorage().Get(auth.LSK_USER_TOKEN, n.token); err != nil {
			slog.Error(fmt.Sprintf("Failed to get user's token from '%s' local storage key.", auth.LSK_USER_TOKEN),
				"error", err)
		}
	} else {
		slog.Info(fmt.Sprintf("User is not logged in ('%s' key was not found in the local storage.", auth.LSK_USER_TOKEN))
	}
	ctx.ObserveState(ShowDropdownState, &n.showDropdown).OnChange(func() {
		ctx.Update()
	})
	n.inited = true
}

// isUserLoggedIn returns true if the user is logged in.
// TODO: and his token(s) are still valid.
func (n *NavUserMenu) isUserLoggedIn() bool {
	return n.token != nil
}

func (n *NavUserMenu) Render() app.UI {

	// If we don't know (yet) whether the user is logged in or not, don't show anything.
	if !n.inited {
		return app.Div().Class("text-sm min-w-[31px]").Body(
			app.Text("."),
		)
	}

	// If the user is not logged in, just show the login link.
	if !n.isUserLoggedIn() {
		return app.A().
			Href("/login").Text("Login").
			Class(navbarLinkCss)
	}
	// Otherwise, show the user icon button.
	return app.Div().
		Class("text-sm text-gray-600 hover:bg-gray-100 rounded-lg transition duration-200 flex flex-col items-end overflow-visible").
		Body(
			app.Button().Class("px-8 py-2 align rounded-lg text-sm outline-none").
				OnClick(func(ctx app.Context, e app.Event) {
					n.showDropdown = !n.showDropdown
					ctx.SetState(ShowDropdownState, n.showDropdown).Persist().Broadcast()
					ctx.Update()
				}).
				Body(
					app.Div().Class("rounded-full justify-center").Body(
						app.Raw(USER_ICON),
					),
				),
			app.If(n.showDropdown, func() app.UI { return &NavUserDropdown{} }),
		)
}

type NavUserDropdown struct {
	app.Compo
}

func (n *NavUserDropdown) Render() app.UI {
	return app.Div().
		Style("width", "100%").Style("height", "1000%").Style("padding", "0").
		Style("position", "absolute").Style("top", "0").Style("left", "0").
		OnClick(func(ctx app.Context, e app.Event) {
			ctx.SetState(ShowDropdownState, false).Persist().Broadcast()
		}).Body(
		app.Div().
			Class("w-20 mt-14 mr-[60px] bg-white rounded-lg shadow-2xl float-right").Body(
			app.Div().Body(
				app.Ul().
					Class("shadow-2xl bg-white py-2 z-[1000] min-w-full w-max rounded-lg max-h-96 overflow-auto").Body(
					app.Li().
						Class(DropdownLinkCss).Body(
						app.A().Href("/my-profile").
							Class("py-2.5 px-5 min-w-full w-max min-h-full flex text-[#333]").Body(
							app.Div().Class("mr-3").Body(
								app.Raw(USER_ICON),
							),
							app.Text("My Profile"),
						),
					),
					// TODO: Show "Data Management" link, if the user has admin rights. Applicable to navbar as well.
					// ...
					app.Li().Class("px-4 py-2").Body(app.Hr()),
					app.Li().Class(DropdownLinkCss).Body(
						app.A().Href("/logout").
							Class("py-2.5 px-5 min-w-full w-max min-h-full flex text-[#333]").Body(
							app.Div().Class("mr-3").Body(
								app.Raw(LOGOUT_ICON),
							),
							app.Text("Logout"),
						),
					),
				),
			),
		),
	)
}
