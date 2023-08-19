package api

func (s *Server) routes() {
	s.router.Use(s.recoverPanic)
}
