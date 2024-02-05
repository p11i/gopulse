package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/p11i/gopulse/pkg/analyzer"
)

var (
	versionFlag bool
	verboseFlag bool
)

func init() {
	flag.BoolVar(&versionFlag, "version", false, "Print the version")
	flag.BoolVar(&verboseFlag, "verbose", false, "Enable verbose mode")
	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Println("GoPulse v1.0.0")
		return
	}

	processesOutput, err := analyzer.GetProcesses()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	analyzer.AnalyzeProcesses(processesOutput)
}
