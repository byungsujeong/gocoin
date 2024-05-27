package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/byungsujeong/gocoin/explorer"
	"github.com/byungsujeong/gocoin/rest"
)

func usage() {
	fmt.Printf("Welcome\n\n")
	fmt.Printf("Please user the following flags:\n\n")
	fmt.Printf("-port:	Set the PORT of the server\n")
	fmt.Printf("-mode:	Choose between 'SSR' and 'API'\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "API", "Choose between 'SSR' and 'API'")

	flag.Parse()

	switch *mode {
	case "API":
		rest.Start(*port)
	case "SSR":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
