package main

import (
    "flag"
    "fmt"
    "os"

    "Containers_Lab/internal/core"
)

func main() {
    mode := flag.String("mode", "C", "C|B|A|Aplus")
    in   := flag.String("in", "", "input file path")
    flag.Parse()
    if *in == "" {
        fmt.Fprintln(os.Stderr, "-in is required")
        os.Exit(2)
    }

    switch *mode {
    case "C":
        // Initialize the Option C static state
		st := core.NewState()
		
		// Load records from the input file path (*in)
		if err := st.LoadFromFile_C(*in); err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}
		
		// Print the resulting sorted job buckets
		st.PrintBuckets()
    case "B":
        // TODO: call bdynamic loader & printer
    case "A":
        // TODO: call adynamic loader & printer with capVehicles=1
    case "Aplus":
        // TODO: call adynamic loader & printer with capVehicles=2 (and extra type)
    default:
        fmt.Fprintln(os.Stderr, "unknown mode")
        os.Exit(2)
    }
}