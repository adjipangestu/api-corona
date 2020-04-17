package controllers

import "corona/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/all", middlewares.SetMiddlewareJSON(s.GetDataNegara)).Methods("GET")
	s.Router.HandleFunc("/provinsi", middlewares.SetMiddlewareJSON(s.GetDataProvinsi)).Methods("GET")
}
