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

// TODO: add happy path case.
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
	// {
	// 	fixturepath: path.Join(fixturesdir, "package.json"),
	// 	expected:    nil,
	// 	err:         nil,
	// },
}
