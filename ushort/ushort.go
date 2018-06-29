package main

import (
	f "fmt"
	flag "github.com/ogier/pflag"
	"os"
)

// flags
var (
	url  string
	urls map[string]int
)

func main() {

	//Parse flags
	flag.Parse()

	// if the user does not supply flags, print usage
	// TODO: put in own function
	if flag.NFlag() == 0 {
		f.Printf("Usage: %s [options]\n", os.Args[0])
		f.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

}

func init() {
	flag.StringVarP(&url, "add", "a", "", "Add shortened url entry")
}
