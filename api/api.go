package api

type Api interface {
	RegisterRoutes()
	Run(port string) error
}
