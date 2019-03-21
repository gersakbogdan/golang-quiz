package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "USAGE\n")
	fmt.Fprintf(os.Stderr, "  quizer <service> [flags]\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Services\n")
	fmt.Fprintf(os.Stderr, "  server        Quizer server\n")
	fmt.Fprintf(os.Stderr, "  client        Quizer CLI client\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Flags\n")
	fmt.Fprintf(os.Stderr, "  -host         Quizer client port.\n")
	fmt.Fprintf(os.Stderr, "  -port         Quizer server/client port.\n")
	fmt.Fprintf(os.Stderr, "  -dbFile       Quizer questions json file\n")
	fmt.Fprintf(os.Stderr, "\n")
}

func main() {
	var (
		port   = flag.Int("port", 3000, "The server port")
		host   = flag.String("host", "localhost", "The server port")
		dbFile = flag.String("dbFile", "data/questions.json", "JSON format questions file")
	)

	flag.Parse()

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "server":
		runServer(*port, *dbFile)
	case "client":
		runClient(*host, *port)
	default:
		usage()
		os.Exit(1)
	}
}
