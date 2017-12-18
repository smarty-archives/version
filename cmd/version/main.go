package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/smartystreets/version"
	"github.com/smartystreets/version/git"
)

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "  cd to a git repository you'd like to version and run one of the following commands:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "  version major\n    (updates 'major' version number: 1.2.3 -> 2.0.0)")
		fmt.Fprintln(os.Stderr, "  version minor\n    (updates 'minor' version number: 1.2.3 -> 1.3.0)")
		fmt.Fprintln(os.Stderr, "  version patch\n    (updates 'patch' version number: 1.2.3 -> 1.2.4)")
		flag.PrintDefaults()
	}
	flag.Parse()

	var repository version.Repository = new(git.Repository)

	previous, err := repository.CurrentVersion()
	if err != nil {
		log.Fatal(err)
	}

	if !previous.Dirty {
		log.Fatalln("No changes since last version:", previous)
		return
	}

	increment := ""
	if args := flag.Args(); len(args) > 0 {
		increment = args[0]
	}

	current := previous.Increment(increment)

	err = repository.UpdateVersion(current)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v -> %v", previous, current)
}
