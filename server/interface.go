package server

// IServer ...
type IServer interface {
	Serve()
	Run()
	Stop()
}
