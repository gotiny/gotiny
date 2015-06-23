package gotiny

import (
	"fmt"
	"net/http"
)


/*
	Base connnection wrapper
	@implements ResponseWriter interface and also wrapes around it
*/

type Connection struct {
	ResponseWriter http.ResponseWriter
	Request  *http.Request
}

func (connection *Connection) WriteString(content string) {
	fmt.Fprintf(connection.ResponseWriter, content)
}

func (connection *Connection) Write(bytes []byte) {
	connection.ResponseWriter.Write(bytes)
}

func (connection *Connection) WriteHeader(header int) {
	connection.ResponseWriter.WriteHeader(header)
}


/*
	A really really simple go http server
	wrapped in WebServer struct
*/


type WebServer struct {
	Addr	string	
	server *http.Server
	mux    *http.ServeMux
}

func (webserver *WebServer) Route(path string, handler func(connection *Connection)) {
	webserver.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		handler(&Connection{Request: request, ResponseWriter: writer})
	})
}

func (webserver *WebServer) Start() {
	webserver.server.ListenAndServe()
}

func (server *TinyServer) Hello () {
	fmt.Println("Hello from WebServer")
}

func NewWebServer(Addr string) *WebServer {
	s := &WebServer{}
	s.Addr = Addr
	s.server = &http.Server{
		Addr: Addr,
	}
	s.mux = http.DefaultServeMux
	handler := new(http.Handler)
	*handler = s.mux
	s.server.Handler = *handler
	return s
}
