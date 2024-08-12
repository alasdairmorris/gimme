package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/docopt/docopt-go"
)

const version = "v0.1"

const usage = `A super-simple file downloader.

Usage:
  gimme -o FILE URL
  gimme -h | --help
  gimme --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.
  -o, --outfile FILE     File to download to.

Homepage: https://github.com/alasdairmorris/gimme
`

type Config struct {
	Outfile string
	Url     string
}

func exitOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// credit: https://stackoverflow.com/a/33853856
func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Parse and validate command-line arguments
func getConfig() Config {

	var (
		retval Config
		opts   docopt.Opts
		err    error
	)

	opts, err = docopt.ParseArgs(usage+" ", nil, version)
	exitOnError(err)

	// Outfile
	retval.Outfile, err = opts.String("--outfile")
	exitOnError(err)

	if _, err = os.Stat(retval.Outfile); err == nil {
		log.Fatal(fmt.Sprintf("Output file %s already exists.", retval.Outfile))
	}

	// URL
	retval.Url, err = opts.String("URL")
	exitOnError(err)

	return retval
}

func main() {
	log.SetFlags(0)
	config := getConfig()

	err := downloadFile(config.Outfile, config.Url)
	if err != nil {
		log.Fatal(err)
	}
}
