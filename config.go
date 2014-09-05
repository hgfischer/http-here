package main

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	Address   string
	CORS      bool
	CacheTime int
}

var fs = &flag.FlagSet{}

func printErrorUsageAndExit(err string, code int) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err)
	printUsage()
	os.Exit(code)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nFlags:\n")
	fs.PrintDefaults()
}

// Parse command line and load config file. The order of precedence, from more important to less important, is:
// cmd line flag > environment variable > config file
func loadConfig() *config {
	cfg := &config{}

	fs.StringVar(&cfg.Address, "a", ":8080", "Address to start listening for HTTP connections")
	fs.IntVar(&cfg.CacheTime, "c", -1, "Set cache time, in seconds, for cache-control max-age header")
	fs.BoolVar(&cfg.CORS, "cors", false, "Enable CORS via the 'Access-Control-Allow-Origin' header")

	fs.Usage = printUsage
	if err := fs.Parse(os.Args[1:]); err != nil {
		printErrorUsageAndExit(err.Error(), 1)
	}
	return cfg
}
