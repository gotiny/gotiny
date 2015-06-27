package gotiny

import (
	"fmt"
	"net/http"
	"net/url"
)

/*
	The GoTiny framework webserver connection
	that uses composition to extend the web.Connection struct
*/

type TinyConnection struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Url            *url.URL
	Vars           map[string]string
}

func (connection *TinyConnection) WriteString(content string) {
	fmt.Fprintf(connection.ResponseWriter, content)
}

func (connection *TinyConnection) Write(bytes []byte) {
	connection.ResponseWriter.Write(bytes)
}

func (connection *TinyConnection) WriteHeader(header int) {
	connection.ResponseWriter.WriteHeader(header)
}

type TinyConnectionHandler func(*TinyConnection)
