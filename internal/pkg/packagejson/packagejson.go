// Package packagejson contains structure to represent package.json
package packagejson

import (
	"encoding/json"
	"errors"
)

// PackageJSON represents NodeJS package.json
type PackageJSON struct {
	Name        string   `json:"name"`    // required
	Version     string   `json:"version"` // required
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
	if len(p.Name) == 0 {
		return errors.New("'name' field is required in package.json")
	}
	return nil
}
