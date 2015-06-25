# gotiny

Tiny, lightweight and barebones web framework written in golang

## Building:

- make build
- make run
- make clean

## Example

Gotiny features really simple and lightweight routing system

    import "gotiny"
    
    // instantiating a GoTiny HTTP server
    tiny := gotiny.NewTinyServer(Addr)
    	
    // adding routes
    tiny.AddHandler("/", func(c *gotiny.TinyConnection){
    	c.WriteString( "<h1>Hey there!</h1>" )
    })
    
    // Routing variables
    // You can access them via c.Vars
    tiny.AddHandler("/page/<page_name>", func(c *gotiny.TinyConnection){
    	c.WriteString( fmt.Sprintf("<h1>This is %s</h1>", c.Vars['page_name']) );
    })

