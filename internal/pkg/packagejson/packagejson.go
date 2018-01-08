// Package packagejson contains structure to represent package.json
package packagejson

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var requiredFields = []string{"Name", "Version"}

// PackageJSON represents NodeJS package.json
type PackageJSON struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Homepage    string   `json:"homepage"`
	License     string   `json:"license"`
	Files       []string `json:"files"`
	Main        string   `json:"main"`
	Os          []string `json:"os"`
	Cpu         []string `json:"cpu"`
	Private     bool     `json:"private"`
}

// Parse parses package.json payload and returns structure.
func Parse(payload []byte) (*PackageJSON, error) {
	var packagejson *PackageJSON
	err := json.Unmarshal(payload, &packagejson)
	return packagejson, err
}

// Validate checks if provided package.json is valid.
func (p *PackageJSON) Validate() error {
	for _, fieldname := range requiredFields {
		value := getField(p, fieldname)
		if len(value) == 0 {
			return fmt.Errorf("'%s' field is required in package.json", fieldname)
		}
	}

	return nil
}

// getField returns struct field value by name.
func getField(i interface{}, fieldname string) string {
	value := reflect.ValueOf(i)
	field := reflect.Indirect(value).FieldByName(fieldname)
	return field.String()
}
