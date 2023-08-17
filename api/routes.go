package api

func (s *Server) routes() {
	s.recoverPanic(s.router)
}
