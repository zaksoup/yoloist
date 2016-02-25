package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
)

type stringSlice []string

func (ss *stringSlice) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *stringSlice) Set(value string) error {
	*ss = append(*ss, value)

	return nil
}

type outputData struct {
	Args []string
	PID  int
}

func main() {
	// store information about this fake process into JSON
	signal.Ignore()

	var data outputData
	data.PID = os.Getpid()
	data.Args = os.Args[1:]

	// validate command line arguments
	// expect them to look like
	//   fake-thing agent -config-dir=/some/path/to/some/dir
	if len(data.Args) == 0 {
		log.Fatal("Expected at least 1 argument")
	}

	// fmt.Fprintf(os.Stdout, "some standard out")
	// fmt.Fprintf(os.Stderr, "some standard error")
	outputBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, string(outputBytes))
}
