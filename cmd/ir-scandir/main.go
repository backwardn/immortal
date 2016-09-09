package main

import (
	"flag"
	"fmt"
	"os"

	ir "github.com/immortal/scandir"
)

var version string

func main() {
	var (
		v    = flag.Bool("v", false, fmt.Sprintf("Print version: %s", version))
		path string
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [dir]\n\n", os.Args[0])
		fmt.Printf("  dir   The directory that will be scanned.\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *v {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if len(flag.Args()) >= 1 {
		path = flag.Args()[0]
	} else {
		flag.Usage()
		os.Exit(1)
	}

	cmd, err := ir.NewScanDir(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Start()
}
