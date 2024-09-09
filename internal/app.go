package app

type App struct {
	Router *Router
}

func New() *App {

	router := InitRouter(":4000")
	return &App{
		Router: router,
	}
}

func (a *App) Run() error {

	a.Router.Run()
	return nil
}
