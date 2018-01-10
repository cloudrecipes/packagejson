package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	p "github.com/cloudrecipes/packagejson/pkg/packagejson"
)

const usage = "\nUsage main.go --path=<PATH_TO_PACKAGE.JSON>"

func errOut(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	filenamePtr := flag.String("path", "", "path to the package.json")
	flag.Parse()
	filename := *filenamePtr

	if len(filename) == 0 {
		fmt.Println("package.json location required")
		fmt.Println(usage)
		os.Exit(1)
	}

	payload, err := ioutil.ReadFile(filename)
	errOut(err)

	packagejson, err := p.Parse(payload)
	errOut(err)

	err = packagejson.Validate()
	errOut(err)

	fmt.Printf("%s is a valid package.json\n", filename)
	os.Exit(0)
}
