package main

import (
	"fmt"
	"io/ioutil"
	"os"

	termutil "github.com/andrew-d/go-termutil"
	"github.com/client9/xson/cson"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inFile = kingpin.Arg("file", "JSON file.").String()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	data, err := readPipeOrFile(*inFile)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	out := cson.ToJSON(data)
	fmt.Print(string(out))
}

// readPipeOrFile reads from stdin if pipe exists, else from provided file
func readPipeOrFile(fileName string) ([]byte, error) {
	if !termutil.Isatty(os.Stdin.Fd()) {
		return ioutil.ReadAll(os.Stdin)
	}
	if fileName == "" {
		return nil, fmt.Errorf("no piped data and no file provided")
	}
	return ioutil.ReadFile(fileName)
}
