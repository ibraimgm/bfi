package main

import (
	"fmt"
	"os"

	"github.com/ibraimgm/bfi/vm"

	getopt "github.com/pborman/getopt/v2"
)

func main() {
	tsFlag := getopt.UintLong("tapesize", 't', 3000, "sets the tape size")
	csFlag := getopt.IntLong("cellsize", 'c', 8, "sets the cell size")
	helpFlag := getopt.BoolLong("help", 'h', "prints this help message")

	if err := getopt.Getopt(nil); err != nil {
		fmt.Printf("%v\n\n", err)
		getopt.Usage()
		os.Exit(1)
	}

	if *helpFlag {
		getopt.Usage()
		return
	}

	if err := vm.CheckCellSize(*csFlag); err != nil {
		fmt.Printf("%v\n\n", err)
		getopt.Usage()
		os.Exit(1)
	}

	args := getopt.Args()
	if len(args) != 1 {
		fmt.Printf("missing file argument\n\n")
		getopt.Usage()
		os.Exit(1)
	}

	bfvm, err := vm.WithSpecs(*csFlag, int(*tsFlag))
	if err != nil {
		fmt.Printf("error creating vm: %v", err)
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("error opening %s: %v", args[0], err)
		os.Exit(1)
	}

	if err := bfvm.LoadFromStream(file); err != nil {
		fmt.Printf("error loading source file: %v", err)
		os.Exit(1)
	}

	if err := bfvm.Run(); err != nil {
		fmt.Printf("error running virtual machine: %v", err)
		os.Exit(1)
	}
}
