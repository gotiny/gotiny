package gotiny

import (
	"net"
	"fmt"
	"runtime"
	"os/exec"
)

var DefaultHost string = "127.0.0.1"

func IsHostAvailable(hostAddr string) (bool, error) {
	
	// Try listening for tcp on a given host

	listener, err := net.Listen("tcp", hostAddr)
	if listener != nil {
		listener.Close()
	}
	
	// If we have no error then the port is available

	var hostAvailable bool = true;
	if err != nil {
		hostAvailable = false

	}

	return hostAvailable, err
}

func IsLocalPortAvailable(port int) (bool, error) {
	
	// If user is only checking for default port
	// then user default host as well

	hostAddr := fmt.Sprintf("%s:%d", DefaultHost, port)

	return IsHostAvailable(hostAddr)
}

func OpenBrowser(url string) {
	// Platform specific browser open thing
	switch(runtime.GOOS) {

		// osx
		case "darwin":
			go func() {
				var cmd *exec.Cmd = exec.Command("open", url)
				cmd.Run()
			}()

		// this is obvious
		case "linux":
			go func() {
				var cmd *exec.Cmd = exec.Command("xdg-open", url)
				cmd.Run()
			}()

		// oh hey windows ... 
		default:
			fmt.Printf("Warning: OpenBrowser() unimplemented for platform '%s'\n", runtime.GOOS )

	}
}