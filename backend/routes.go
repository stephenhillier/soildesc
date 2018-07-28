package main

func (s *Server) routes() {
	s.router.Post("/describe", s.Describe)
}
