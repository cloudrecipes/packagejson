package packagejson_test

import (
	"io/ioutil"
	"reflect"
	"testing"

	p "github.com/cloudrecipes/packagejson"
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

func TestParse(t *testing.T) {
	for _, test := range parseTestCases {
		payload, err := ioutil.ReadFile(test.fixturepath)
		if err != nil {
			t.Fatalf("\n>>> Expected err to be nil, but got:\n%v", err)
		}

		actual, err := p.Parse(payload)
		if test.err != nil {
			if err == nil || test.err.Error() != err.Error() {
				t.Fatalf("\n>>> Expected validation error:\n%v\n<<< but got:\n%v", test.err, err)
			}
			continue
		}

		if test.err == nil && err != nil {
			t.Fatalf("\n>>> Expected validation error:\nnil\n<<< but got:\n%v", err)
		}

		if _, equals := actual.Equals(test.expected); !equals {
			t.Fatalf("\n>>> Expected:\n%v\n<<< but got:\n%v", test.expected, actual)
		}
	}
}

func TestEquals(t *testing.T) {
	for _, test := range equalsTestCases {
		diff, equals := test.this.Equals(test.other)

		if equals != test.equals {
			t.Fatalf("\n>>> Expected:\n%t\n<<< but got:\n%t", test.equals, equals)
		}

		if !reflect.DeepEqual(diff, test.diff) {
			t.Fatalf("\n>>> Expected:\n%v\n<<< but got:\n%v", test.diff, diff)
		}
	}
}

func TestTest(t *testing.T) {
	expected := "cross-env NODE_ENV=test nyc --all ava"
	actual := packageJson.Test()
	if actual != expected {
		t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", expected, actual)
	}
}
