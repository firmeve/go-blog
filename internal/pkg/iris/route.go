package iris

type Route interface {
	registerRoutes()
	registerMacros()
}