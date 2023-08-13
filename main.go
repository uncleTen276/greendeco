package main

import (
	"github.com/sekke276/greendeco.git/cmd/server"
)

// @title Fiber Go API
// @version 1.0
// @description greendeco
// @contact.name Nguyen Tri
// @contact.email tringuyen2762001@gmail.com
// @termsOfService
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	server.Serve()
}
