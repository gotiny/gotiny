package gotiny

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

/*
	A custom TCP listener, that supports
	stopping of the server
*/

type TinyTCPListener struct {
	*net.TCPListener
	stop chan int
}

func NewTinyTCPListener(listener net.Listener) (*TinyTCPListener, error) {
	tcpListener, canWrapListener := listener.(*net.TCPListener)
	if !canWrapListener {
		return nil, errors.New("Cannot wrap listener")
	}

	ttcpl := &TinyTCPListener{}
	ttcpl.TCPListener = tcpListener
	ttcpl.stop = make(chan int)

	return ttcpl, nil
}

func (listener *TinyTCPListener) Accept() (net.Conn, error) {
	for {
		listener.TCPListener.SetDeadline(time.Now().Add(time.Second))

		conn, err := listener.TCPListener.Accept()

		fmt.Println("> Accept()")

		select {
		case <-listener.stop:
			// Channel has been closed.
			// Stop processing requests
			return nil, errors.New("Server was stopped")
		default:
			// Channel still open
		}

		if err != nil {
			netErr, ok := err.(net.Error)

			// Verify if error is caused by timeout
			// and not due to any other error raised in Accept()
			if ok && netErr.Timeout() && netErr.Temporary() {
				continue
			}
		}

		// If any other err, or success, return them
		return conn, err
	}
}

/*
	A really really simple go http server
	that also supports cool routing
*/

type TinyServer struct {
	Addr   string
	Server *http.Server
	Mux    *http.ServeMux

	Listener *TinyTCPListener
	Waiter   sync.WaitGroup

	Routes   []*Route
	Handlers []TinyConnectionHandler
}

func (tiny *TinyServer) AddRouteHandler(route *Route, routeHandler TinyConnectionHandler) {
	tiny.Routes = append(tiny.Routes, route)
	tiny.Handlers = append(tiny.Handlers, routeHandler)
}

func (tiny *TinyServer) AddHandler(format string, routeHandler TinyConnectionHandler) {
	route := NewRoute(format)
	tiny.Routes = append(tiny.Routes, route)
	tiny.Handlers = append(tiny.Handlers, routeHandler)
}

func (tiny *TinyServer) DefaultHandler(connection *TinyConnection) {
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

	fmt.Println("Starting server at : ", tiny.Server.Addr)

	defaultListener, err := net.Listen("tcp", tiny.Server.Addr)
	if err == nil {
		// Dispatch the server to another go routine
		// go func() {
		tiny.Waiter.Add(1)
		defer tiny.Waiter.Done()

		fmt.Println("> NewTinyTCPListener()")
		tiny.Listener, _ = NewTinyTCPListener(defaultListener)
		tiny.Server.Serve(tiny.Listener)
		// fmt.Println("> listener ended")
		// }();
	} else {
		panic(err)
	}
}

func (tiny *TinyServer) Stop() {
	close(tiny.Listener.stop)
}

func (tiny *TinyServer) Waitup() {
	tiny.Waiter.Wait()
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
		tiny.DefaultHandler(&TinyConnection{ResponseWriter: writer, Request: request})
	})

	// Setup public defaults
	tiny.Routes = make([]*Route, 0)
	tiny.Handlers = make([]TinyConnectionHandler, 0)

	return tiny
}
