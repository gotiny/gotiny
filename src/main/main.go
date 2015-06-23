package main

import (
	"fmt"
	"gotiny"
)

func main() {
	
	// host running the Gotiny HTTP server instance
	Addr := "localhost:8081"

	// Addr checking
	if _, error := gotiny.IsHostAvailable(Addr); error != nil {
		panic(error)
	}

	
	// instantiating a GoTiny HTTP server
	tiny := gotiny.NewTinyServer(Addr)
	
	// adding routes
	tiny.AddHandler("/", func(c *gotiny.TinyConnection){
		c.WriteString( "<h1>Hey there!</h1>" )
	})
	tiny.AddHandler("/page/about", func(c *gotiny.TinyConnection){
		c.WriteString( "Special about page!" )
	})
	tiny.AddHandler("/page/<page_id>", func(c *gotiny.TinyConnection){
		c.WriteString( fmt.Sprint(c.Vars, "\n") )
	})
	tiny.AddHandler("/page/<page_id>/info", func(c *gotiny.TinyConnection){
		c.WriteString( fmt.Sprint("Student Info:", c.Vars, "\n") )
	})


	gotiny.OpenBrowser("http://localhost:8081")

	// go func() {
	// 	var cmd *exec.Cmd = exec.Command("open", HTTPAddr)
	// 	cmd.Run()
	// }()



	// start the server. Woohoo!
	tiny.Start()


}
