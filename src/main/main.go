package main

import (
	"fmt"
	"os/exec"
	"gotiny"
	// "html/template"
	// "bytes"
)

func main() {
	fmt.Println("Hi")

	Addr := ":8080"
	HostURL := fmt.Sprintf("http://localhost%s", Addr)

	go func() {
		var cmd *exec.Cmd = exec.Command("open", HostURL)
		cmd.Run()
	}()

	// server := web.NewWebServer()

	// server.Route("/", func(connection *web.Connection) {
	// 	tmpl, _ := template.New("index.tmpl").Parse("Hello {{.Name}} {{.Salute}}!")
	// 	var doc bytes.Buffer
	// 	vars := make(map[string]string)
	// 	vars["Name"] = "Kalyan!"
	// 	vars["Salute"] = "Good Morning!"
	// 	tmpl.Execute( &doc, vars )
	// 	connection.Write(doc.String())
	// })
	// server.Route("/test/?", func(connection *web.Connection) {
	// 	connection.Write("Test Yeah!")
	// })

	// server.Start()

	tiny := gotiny.NewTinyServer(Addr)
	tiny.AddHandler("/", func(c *gotiny.TinyConnection){
		c.Write( "Home" )
	})
	tiny.AddHandler("/page/<page_id>", func(c *gotiny.TinyConnection){
		c.Write( fmt.Sprint(c.Vars, "\n") )
	})
	tiny.AddHandler("/page/<page_id>/info", func(c *gotiny.TinyConnection){
		c.Write( fmt.Sprint("Student Info:", c.Vars, "\n") )
	})
	tiny.Start()
}
