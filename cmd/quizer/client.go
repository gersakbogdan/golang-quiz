package main

import (
	"github.com/gersakbogdan/fasttrackquiz/services/client"
)

func runClient(host string, port int) {
	client.Run(host, port)
}
