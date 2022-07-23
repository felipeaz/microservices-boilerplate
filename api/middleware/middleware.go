package middleware

type Middleware interface {
	Middleware()
}

func New() Middleware {
	return &middleware{}
}

type middleware struct {
}

func (m *middleware) Middleware() {

}
