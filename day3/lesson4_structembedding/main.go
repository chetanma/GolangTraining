package main

//Struct embedding

type Server struct {
	Name string
	*Logger
}

func (s *Server) Start() {
	//s.Logger.Info("server  started")
	s.Info("server has started")
}

func main() {
	logger := NewLogger()

	server := Server{
		Name:   "Server-A",
		Logger: &logger,
	}

	server.Info("Starting the server")
	server.Start()
	server.Debug("INvoking ops 13A1")
	server.Debug("INvoking ops 12A1")
	server.DumpToConsole()

}

func JustUseSimpleLogger() {
	logger := NewLogger()

	logger.Info("In main")
	/// imagine invoking some funcs

	logger.Info("Invoking step 1")
	logger.Debug("Val of var TDDF is 2323")
	logger.Info("Invoking step 2")
	// soem calls
	logger.Debug("Val of var RRR is ddd")

	logger.DumpToConsole()

}
