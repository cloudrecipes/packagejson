package packagejson_test

import (
	"errors"
	"os"
	"path"

	p "github.com/cloudrecipes/package-json/internal/pkg/packagejson"
)

var fixturesdir = path.Join(os.Getenv("GOPATH"), "src", "github.com",
	"cloudrecipes", "package-json", "test", "fixtures")

var validationTestCases = []struct {
	fixturepath string
	err         error
}{
	{fixturepath: path.Join(fixturesdir, "package_no_name.json"), err: errors.New("'Name' field is required in package.json")},
	{fixturepath: path.Join(fixturesdir, "package_no_version.json"), err: errors.New("'Version' field is required in package.json")},
	{fixturepath: path.Join(fixturesdir, "package.json"), err: nil},
}

var packageJson = &p.PackageJSON{
	Name:        "seed",
	Version:     "0.0.1",
	Description: "NodeJS projects seed",
	Keywords:    []string{"nodejs", "seed"},
	Homepage:    "",
	License:     "MIT",
	Files:       []string{},
	Main:        "index.js",
	Scripts: map[string]string{
		"test": "cross-env NODE_ENV=test nyc --all ava",
	},
	Os:      []string{},
	Cpu:     []string{},
	Private: false,
}

var parseTestCases = []struct {
	fixturepath string
	expected    *p.PackageJSON
	err         error
}{
	{
		fixturepath: path.Join(fixturesdir, "broken_package.json"),
		expected:    nil,
		err:         errors.New("unexpected end of JSON input"),
	},
	{
		fixturepath: path.Join(fixturesdir, "package.json"),
		expected:    packageJson,
		err:         nil,
	},
}

var equalsTestCases = []struct {
	this   *p.PackageJSON
	other  *p.PackageJSON
	diff   []string
	equals bool
}{
	{
		this: packageJson,
		other: &p.PackageJSON{
			Name:        "Seed",
			Version:     "0.0.0",
			Description: "NodeJS projects seed!!!",
			Keywords:    []string{"nodejs", ""},
			Homepage:    "Home",
			License:     "MIT2",
			Files:       []string{"main.go"},
			Main:        "index.js",
			Scripts: map[string]string{
				"test": "cross-env NODE_ENV=test nyc --all ava",
			},
			Os:      []string{"osx"},
			Cpu:     []string{"darwin"},
			Private: true,
		},
		diff: []string{
			"Difference @Name: 'seed' vs 'Seed'",
			"Difference @Version: '0.0.1' vs '0.0.0'",
			"Difference @Description: 'NodeJS projects seed' vs 'NodeJS projects seed!!!'",
			"Difference @Keywords: [nodejs seed] vs [nodejs ]",
			"Difference @Homepage: '' vs 'Home'",
			"Difference @License: 'MIT' vs 'MIT2'",
			"Difference @Files: [] vs [main.go]",
			"Difference @Os: [] vs [osx]",
			"Difference @Cpu: [] vs [darwin]",
			"Difference @Private: 'false' vs 'true'",
		},
		equals: false,
	},
	{
		this: packageJson,
		other: &p.PackageJSON{
			Name:        "seed",
			Version:     "0.0.1",
			Description: "NodeJS projects seed",
			Keywords:    []string{"nodejs"},
			Homepage:    "",
			License:     "MIT",
			Files:       []string{},
			Main:        "index.js",
			Scripts:     map[string]string{},
			Os:          []string{},
			Cpu:         []string{},
			Private:     false,
		},
		diff: []string{
			"Difference @Keywords: [nodejs seed] vs [nodejs]",
			"Difference @Scripts: map[test:cross-env NODE_ENV=test nyc --all ava] vs map[]",
		},
		equals: false,
	},
	{
		this: &p.PackageJSON{
			Name:        "seed",
			Version:     "0.0.1",
			Description: "NodeJS projects seed",
			Keywords:    []string{},
			Homepage:    "",
			License:     "MIT",
			Files:       []string{"test.go"},
			Main:        "index.js",
			Scripts:     map[string]string{},
			Os:          []string{"osx"},
			Cpu:         []string{"darwin"},
			Private:     false,
		},
		other: &p.PackageJSON{
			Name:        "seed",
			Version:     "0.0.1",
			Description: "NodeJS projects seed",
			Keywords:    []string{},
			Homepage:    "",
			License:     "MIT",
			Files:       []string{"main.go"},
			Main:        "main.js",
			Scripts:     map[string]string{},
			Os:          []string{"win"},
			Cpu:         []string{"solaris"},
			Private:     false,
		},
		diff: []string{
			"Difference @Files: [test.go] vs [main.go]",
			"Difference @Main: 'index.js' vs 'main.js'",
			"Difference @Os: [osx] vs [win]",
			"Difference @Cpu: [darwin] vs [solaris]",
		},
		equals: false,
	},
}
