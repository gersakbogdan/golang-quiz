package main

import (
	"github.com/gersakbogdan/fasttrackquiz/services/quizer"
)

func runServer(port int, dbFile string) error {
	return quizer.NewServer(dbFile).Run(port)
}
