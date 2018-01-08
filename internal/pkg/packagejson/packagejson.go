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
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Keywords    []string          `json:"keywords"`
	Homepage    string            `json:"homepage"`
	License     string            `json:"license"`
	Files       []string          `json:"files"`
	Main        string            `json:"main"`
	Scripts     map[string]string `json:"scripts"`
	Os          []string          `json:"os"`
	Cpu         []string          `json:"cpu"`
	Private     bool              `json:"private"`
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

// Equals checks if two structures are equal by value.
func (p *PackageJSON) Equals(other *PackageJSON) (diff []string, equals bool) {
	if p.Name != other.Name {
		diff = append(diff, fmt.Sprintf("Difference @Name: '%s' vs '%s'", p.Name, other.Name))
	}

	if p.Version != other.Version {
		diff = append(diff, fmt.Sprintf("Difference @Version: '%s' vs '%s'", p.Version, other.Version))
	}

	if p.Description != other.Description {
		diff = append(diff, fmt.Sprintf("Difference @Description: '%s' vs '%s'", p.Description, other.Description))
	}

	if len(p.Keywords) == 0 && len(other.Keywords) == 0 {
	} else if len(p.Keywords) != len(other.Keywords) {
		diff = append(diff, fmt.Sprintf("Difference @Keywords: %v vs %v", p.Keywords, other.Keywords))
	} else if !reflect.DeepEqual(p.Keywords, other.Keywords) {
		diff = append(diff, fmt.Sprintf("Difference @Keywords: %v vs %v", p.Keywords, other.Keywords))
	}

	if p.Homepage != other.Homepage {
		diff = append(diff, fmt.Sprintf("Difference @Homepage: '%s' vs '%s'", p.Homepage, other.Homepage))
	}

	if p.License != other.License {
		diff = append(diff, fmt.Sprintf("Difference @License: '%s' vs '%s'", p.License, other.License))
	}

	if len(p.Files) == 0 && len(other.Files) == 0 {
	} else if len(p.Files) != len(other.Files) {
		diff = append(diff, fmt.Sprintf("Difference @Files: %v vs %v", p.Files, other.Files))
	} else if !reflect.DeepEqual(p.Files, other.Files) {
		diff = append(diff, fmt.Sprintf("Difference @Files: %v vs %v", p.Files, other.Files))
	}

	if p.Main != other.Main {
		diff = append(diff, fmt.Sprintf("Difference @Main: '%s' vs '%s'", p.Main, other.Main))
	}

	if !reflect.DeepEqual(p.Scripts, other.Scripts) {
		diff = append(diff, fmt.Sprintf("Difference @Scripts: %v vs %v", p.Scripts, other.Scripts))
	}

	if len(p.Os) == 0 && len(other.Os) == 0 {
	} else if len(p.Os) != len(other.Os) {
		diff = append(diff, fmt.Sprintf("Difference @Os: %v vs %v", p.Os, other.Os))
	} else if !reflect.DeepEqual(p.Os, other.Os) {
		diff = append(diff, fmt.Sprintf("Difference @Os: %v vs %v", p.Os, other.Os))
	}

	if len(p.Cpu) == 0 && len(other.Cpu) == 0 {
	} else if len(p.Cpu) != len(other.Cpu) {
		diff = append(diff, fmt.Sprintf("Difference @Cpu: %v vs %v", p.Cpu, other.Cpu))
	} else if !reflect.DeepEqual(p.Cpu, other.Cpu) {
		diff = append(diff, fmt.Sprintf("Difference @Cpu: %v vs %v", p.Cpu, other.Cpu))
	}

	if p.Private != other.Private {
		diff = append(diff, fmt.Sprintf("Difference @Private: '%t' vs '%t'", p.Private, other.Private))
	}

	return diff, len(diff) == 0
}

// getField returns struct field value by name.
func getField(i interface{}, fieldname string) string {
	value := reflect.ValueOf(i)
	field := reflect.Indirect(value).FieldByName(fieldname)
	return field.String()
}
