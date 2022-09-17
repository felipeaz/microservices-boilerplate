package api

type Api interface {
	RegisterRoutes()
	GetRouter() Router
}
