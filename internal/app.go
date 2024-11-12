package app

import "github.com/laply/study-golang-coin/explorer"

type App struct {
	Router *explorer.Router
}

func New() *App {

	router := explorer.InitRouter(":4000")

	return &App{
		Router: router,
	}
}

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

func (a *App) Run() {
	a.Router.Run(port, templateDir)
}
