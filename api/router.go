package api

type Router interface {
	Run(addr ...string) error
}
