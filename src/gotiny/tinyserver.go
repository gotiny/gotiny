package gotiny

import (
	"fmt"
	"net/url"
)


/*
	The GoTiny framework webserver connection
	that uses composition to extend the web.Connection struct
*/

type TinyConnection struct {
	Connection
	Url *url.URL
	Vars map[string]string
}

type TinyConnectionHandler func(*TinyConnection)

/*
	The Server that also supports cool routing
*/

type TinyServer struct {
	server *WebServer
	routes []*Route
	handlers []TinyConnectionHandler
}

func (tiny *TinyServer) AddRouteHandler (route *Route, routeHandler TinyConnectionHandler) {
	tiny.routes = append( tiny.routes, route )
	tiny.handlers = append(tiny.handlers, routeHandler)
}

func (tiny *TinyServer) AddHandler (format string, routeHandler TinyConnectionHandler) {
	route := NewRoute(format)
	tiny.routes = append( tiny.routes, route )
	tiny.handlers = append(tiny.handlers, routeHandler)
}

func (tiny *TinyServer) DefaultHandler (connection *Connection) {
	path := connection.Request.URL.Path

	// Identify which Route is applicable
	for i := range tiny.handlers {
		route := tiny.routes[i]
		matches := route.Match(path)
		if matches != nil {
			conn := &TinyConnection{}
			conn.Request = connection.Request
			conn.ResponseWriter = connection.ResponseWriter
			conn.Url = connection.Request.URL
			conn.Vars = matches

			handler := tiny.handlers[i]
			handler(conn)
			break
		}
	}
}

func (tiny *TinyServer) Start() {
	
	fmt.Println("Starting server at : ", tiny.server.Addr);
	
	tiny.server.Start()
}

func NewTinyServer(Addr string) *TinyServer {

	fmt.Println("### GoTiny says Hello! ###")
	tiny := &TinyServer{}
	tiny.server = NewWebServer(Addr)
	tiny.routes = make([]*Route,0)
	tiny.handlers = make([]TinyConnectionHandler,0)
	tiny.server.Route("/", tiny.DefaultHandler);

	return tiny
}
