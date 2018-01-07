package packagejson_test

import (
	"io/ioutil"
	"testing"

	p "github.com/cloudrecipes/package-json/internal/pkg/packagejson"
)

func TestValidate(t *testing.T) {
	for _, test := range validationTestCases {
		payload, err := ioutil.ReadFile(test.fixturepath)
		if err != nil {
			t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
		}

		packagejson, err := p.Parse(payload)
		if err != nil {
			t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
		}

		err = packagejson.Validate()
		if test.err != nil {
			if err == nil || test.err.Error() != err.Error() {
				t.Fatalf("\n>>> Expected validation error:\n%v\n<<< but got:\n%v", test.err, err)
			}
			continue
		}

		if test.err == nil && err != nil {
			t.Fatalf("\n>>> Expected validation error:\nnil\n<<< but got:\n%v", err)
		}
	}
}
