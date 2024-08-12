package main

import (
	"fmt"
	"strconv"

	"github.com/docopt/docopt-go"
)

const version = "v0.1"

const usage = `A super-simple file downloader.

Usage:
  gimme -p PORT
  gimme -h | --help
  gimme --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.
  -p, --port PORT        Port to listen on.

Homepage: https://github.com/alasdairmorris/gimme
`

type Config struct {
	Port int
}

func exitOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// Parse and validate command-line arguments
func getConfig() Config {

	var (
		retval Config
		opts   docopt.Opts
		port   string
		err    error
	)

	opts, err = docopt.ParseArgs(usage+" ", nil, version)
	exitOnError(err)

	// Port
	port, err = opts.String("--port")
	exitOnError(err)

	retval.Port, err = strconv.Atoi(port)
	exitOnError(err)

	return retval
}

func main() {
	var config = getConfig()
	fmt.Println(config)
}
