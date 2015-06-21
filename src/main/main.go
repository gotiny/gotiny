package main

import (
	"fmt"
	"os/exec"
	"gotiny"
	"net"
)

func main() {
	
	// port running the Gotiny HTTP server instance
	portAddr := ":8081"

	// Addr checking
	listener, err := net.Listen("tcp", portAddr)
	if err != nil {
		panic(err)
	}
	listener.Close()

	// host URL running the instance
	hostUrl := fmt.Sprintf( "http://localhost%s", portAddr )
	
	// instantiating a GoTiny HTTP server
	tiny := gotiny.NewTinyServer(hostUrl, portAddr)
	
	// adding routes
	tiny.AddHandler("/", func(c *gotiny.TinyConnection){
		c.Write( "<h1>Hey there!</h1>" )
	})
	tiny.AddHandler("/page/<page_id>", func(c *gotiny.TinyConnection){
		c.Write( fmt.Sprint(c.Vars, "\n") )
	})
	tiny.AddHandler("/page/<page_id>/info", func(c *gotiny.TinyConnection){
		c.Write( fmt.Sprint("Student Info:", c.Vars, "\n") )
	})

	// start the server. Woohoo!
	tiny.Start()

	go func() {
		var cmd *exec.Cmd = exec.Command("open", hostUrl)
		cmd.Run()
	}()
}
