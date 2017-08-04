package server

type Server interface {
	Serve(hst string, prt int)
}
