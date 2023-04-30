package ports

type Server interface {
	Serve(port int)
}
