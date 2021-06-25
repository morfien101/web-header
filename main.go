package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
)

var (
	version = "0.0.0"

	flagHelp    = flag.Bool("h", false, "Show the help menu")
	flagVersion = flag.Bool("v", false, "Show version number for web-headers")
	flagPort    = flag.Int("p", 9000, "Port number to run on.")
	flagHost    = flag.String("H", "127.0.0.1", "host IP to bind the service to. Use 0.0.0.0 for all.")
)

func main() {
	flag.Parse()

	if *flagHelp {
		flag.PrintDefaults()
		return
	}

	if *flagVersion {
		fmt.Printf("v%s", version)
		return
	}

	hostPort := fmt.Sprintf("%s:%d", *flagHost, *flagPort)
	fmt.Printf("Starting http server on: %s\n", hostPort)

	http.HandleFunc("/", helloHeaders)
	http.ListenAndServe(hostPort, nil)
}

func helloHeaders(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(requestDump))
}
