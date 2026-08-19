package environment

import (
	"fmt"
	"strings"
)

func Load(serverAddress, httpPort, token string) *Configuration {
	return &Configuration{
		Connection: Connection{
			Token:           fmt.Sprint("jwt=", token),
			ContextDeadline: 5,
		},
		Server: Server{
			Address:  serverAddress,
			HttpPort: httpPort,
			Url:      newServerUrl(serverAddress, httpPort),
		},
	}
}

func newServerUrl(serverAddress, serverPort string) string {
	addr := serverAddress
	if !strings.HasPrefix(addr, "http://") && !strings.HasPrefix(addr, "https://") {
		addr = "http://" + addr
	}
	addr = strings.TrimRight(addr, "/")

	if len(strings.TrimSpace(serverPort)) == 0 {
		return fmt.Sprintf("%s/", addr)
	}
	return fmt.Sprintf("%s:%s/", addr, serverPort)
}
