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
	{fixturepath: path.Join(fixturesdir, "package_no_name.json"), err: errors.New("'name' field is required in package.json")},
	{fixturepath: path.Join(fixturesdir, "package.json"), err: nil},
}

// TODO: it's a placeholder for parser validation.
// Expected field should be filled with the correct structure.
var parseTestCases = []struct {
	fixturepath string
	expected    *p.PackageJSON
	err         error
}{
	{
		fixturepath: path.Join(fixturesdir, "package.json"),
		expected:    nil,
		err:         nil,
	},
}
