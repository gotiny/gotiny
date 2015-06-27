package gotiny

import (
	"fmt"
	"net/http"
)

/*
	A really really simple go http server
	that also supports cool routing
*/

type TinyServer struct {
	Addr	string	
	Server *http.Server
	Mux    *http.ServeMux

	Routes []*Route
	Handlers []TinyConnectionHandler
}

func (tiny *TinyServer) AddRouteHandler (route *Route, routeHandler TinyConnectionHandler) {
	tiny.Routes = append( tiny.Routes, route )
	tiny.Handlers = append(tiny.Handlers, routeHandler)
}

func (tiny *TinyServer) AddHandler (format string, routeHandler TinyConnectionHandler) {
	route := NewRoute(format)
	tiny.Routes = append( tiny.Routes, route )
	tiny.Handlers = append(tiny.Handlers, routeHandler)
}

func (tiny *TinyServer) DefaultHandler (connection *TinyConnection) {
	path := connection.Request.URL.Path

	// Identify which Route is applicable
	for i := range tiny.Handlers {
		route := tiny.Routes[i]
		matches := route.Match(path)
		if matches != nil {
			conn := &TinyConnection{}
			conn.Request = connection.Request
			conn.ResponseWriter = connection.ResponseWriter
			conn.Url = connection.Request.URL
			conn.Vars = matches

			handler := tiny.Handlers[i]
			handler(conn)
			break
		}
	}
}

func (tiny *TinyServer) Start() {
	
	fmt.Println("Starting server at : ", tiny.Server.Addr);
	
	tiny.Server.ListenAndServe()
}

func NewTinyServer(Addr string) *TinyServer {

	fmt.Println("### GoTiny says Hello! ###")

	// Setup server
	tiny := &TinyServer{}
	tiny.Server = &http.Server{
		Addr: Addr,
	}

	// Create default muxer
	// Todo: refactor
	tiny.Mux = http.DefaultServeMux
	handler := new(http.Handler)
	*handler = tiny.Mux
	tiny.Server.Handler = *handler

	tiny.Mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			tiny.DefaultHandler(&TinyConnection{ResponseWriter:writer, Request:request})
	})

	// Setup public defaults
	tiny.Routes = make([]*Route,0)
	tiny.Handlers = make([]TinyConnectionHandler,0)

	return tiny
}
